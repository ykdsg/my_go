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
}
type Task func()

func (p *Pool) run() {
	idx := 0
	for {
		select {
		//监听quit，如果有信号就退出
		case <-p.quit:
			return
		//如果能放入active，说明还有空间
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

var ErrWorkerPoolFreed = errors.New("workerpool1 freed")

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	//一旦 p.tasks 可写，提交的 Task 就会被写入 tasks channel，worker线程就会收到channel 传递过来的task
	case p.tasks <- t:
		return nil

	}

}

func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Printf("workpool freed\n")
}

func New(capacity int) *Pool {
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
		//capacity 决定了初始化的worker数量
		active: make(chan struct{}, capacity),
	}
	fmt.Printf("workerpool1 start\n")

	go p.run()
	return p

}
