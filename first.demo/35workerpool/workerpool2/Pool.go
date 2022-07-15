package workerpool

import (
	"errors"
	"fmt"
	"sync"
)

const (
	defaultCapacity = 100
	maxCapacity     = 1000
)

type Pool struct {
	capacity int           // workerpool大小
	active   chan struct{} // 对应上图中的active channel ，有缓冲区channel
	tasks    chan Task     // 对应上图中的task channel

	wg   sync.WaitGroup // 用于在pool销毁时等待所有worker退出
	quit chan struct{}  // 用于通知各个worker退出的信号channel

	preAllo bool //是否在创建pool的时候就预创建workers，默认值为：false
	// 当pool满的情况下，新的Schedule调用是否阻塞当前goroutine。默认值：true
	// 如果block = false，则Schedule返回ErrNoWorkerAvailInPool
	block bool
}
type Task func()

func (p *Pool) run() {
	idx := len(p.active)

	//如果没有初始化
	if !p.preAllo {
	loop:
		//刚开始tasks 是空的，所以会在这里阻塞
		for t := range p.tasks {
			//这里的task仅是触发了worker创建，这里是调度循环，不处理task，所以 要把task扔回tasks channel，等worker启动后再处理。
			p.returnTaks(t)
			select {
			case <-p.quit:
				return
			case p.active <- struct{}{}:
				idx++
				println("")
				p.newWorker(idx)
			default:
				break loop

			}
		}
	}

	for {
		select {
		case <-p.quit:
			return
		case p.active <- struct{}{}:
			idx++
			p.newWorker(idx)
		}
	}
}

func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover panic[%s] and exit\n", i, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d]: start\n", i)
		for true {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d]:exit\n", i)
				<-p.active
				return
			case t := <-p.tasks:
				fmt.Printf("worker[%03d]: receive a task\n", i)
				t()
			}

		}
	}()

}

var ErrWorkerPoolFreed = errors.New("workerpool freed")

var ErrNoIdleWorkerInPool = errors.New("no idle worker in pool")

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	//一旦 p.tasks 可写，提交的 Task 就会被写入 tasks channel，worker线程就会收到channel 传递过来的task
	case p.tasks <- t:
		return nil
	default:
		//根据 block 字段的值，决定究竟是继续阻塞在 tasks channel 上，还是返回 ErrNoIdleWorkerInPool 错误。
		if p.block {
			p.tasks <- t
			return nil
		}
		return ErrNoIdleWorkerInPool

	}

}

func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Printf("workpool freed\n")
}

func (p *Pool) returnTaks(t Task) {
	go func() {
		p.tasks <- t
	}()
}

func New(capacity int, opts ...Option) *Pool {

	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	p := &Pool{
		capacity: capacity,
		//tasks 是一个无缓冲的 channel，如果 pool 中 worker 数量已达上限，而且 worker 都在处理 task 的状态，
		//那么 Schedule 方法就会阻 塞，直到有 worker 变为 idle 状态来读取 tasks channel
		tasks: make(chan Task),
		quit:  make(chan struct{}),
		//capacity 决定了worker数量
		active: make(chan struct{}, capacity),
	}

	for _, opt := range opts {
		opt(p)

	}
	fmt.Printf("workerpool start(preAlloc=%t)\n", p.preAllo)

	if p.preAllo {
		// create all goroutines and send into works channel
		for i := 0; i < p.capacity; i++ {
			p.newWorker(i + 1)
			p.active <- struct{}{}
		}
	}

	go p.run()
	return p

}
