package workpool1

import (
	"errors"
	"fmt"
	"sync"
)

const (
	defaultCapacity = 100
	maxCapacity     = 10000
)

type Pool struct {
	// worker 的最大容量
	capacity int
	// worker 的计数器
	active chan struct{}
	tasks  chan Task
	// 用于通知各个worker退出
	quite chan struct{}
	wg    sync.WaitGroup // 用于在pool销毁时等待所有worker退出
}

type Task func()

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quite:
		return errors.New("workerpool1 freed")
	// 	一旦tasks 可写
	case p.tasks <- t:
		return nil

	}
}

func (p *Pool) Free() {
	close(p.quite)
	p.wg.Wait()
	fmt.Println("workpool freed")
}

func (p *Pool) run() {
	idx := 0
	for {
		select {
		case <-p.quite:
			return
		case p.active <- struct{}{}:
			idx++
			p.newWorker(idx)
		}
	}

}

func (p *Pool) newWorker(idx int) {
	fmt.Printf("worker[%03d]: start\n", idx)
	p.wg.Add(1)
	go func() {
		fmt.Printf("worker[%03d]: start+++++\n", idx)
		defer func() {
			// 为了防止用户提交的 task 抛出 panic,进而导致整个 workerpool 受到影响
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover panic[%s] and exit\n", idx, err)
				// 更新 worker 计数器
				<-p.active
			}
			p.wg.Done()
		}()
		for {
			select {
			case <-p.quite:
				<-p.active
				fmt.Printf("worker[%03d]: quite-----\n", idx)
				return
			case task := <-p.tasks:
				fmt.Printf("worker[%03d]: receive a task\n", idx)
				task()
			}

		}
	}()
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
		active:   make(chan struct{}, capacity),
		tasks:    make(chan Task),
		quite:    make(chan struct{}),
	}

	fmt.Println("workerpool1 start+++++")
	go p.run()
	return p
}
