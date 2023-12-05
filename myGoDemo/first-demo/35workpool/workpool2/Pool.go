package workpool2

import (
	"errors"
	"fmt"
	"sync"
)

type Task func()

type Pool struct {
	capacity int
	active   chan struct{}
	tasks    chan Task
	quite    chan struct{}
	wg       sync.WaitGroup // WaitGroup 是零值可用的
}

func (p *Pool) run() {
	idx := 0

	for {
		select {
		case <-p.quite:
			return
		case p.active <- struct{}{}:
			p.newWorker(idx)
			idx++
		}

	}
}

func (p *Pool) newWorker(idx int) {
	p.wg.Add(1)
	go func() {
		fmt.Printf("worker[%03d]: start+++++\n", idx)
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover panic[%s] and exit\n", idx, err)
				<-p.active
			}
			fmt.Println("defer**************")
			p.wg.Done()
		}()
		for {
			select {
			case <-p.quite:
				p.wg.Done()
				return
			case task := <-p.tasks:
				fmt.Printf("worker[%03d]: receive a task\n", idx)
				task()
			}
		}
	}()

}

func (p *Pool) Free() {
	close(p.quite)
	p.wg.Wait()
	fmt.Println("workpool freed @@@@@@@@@@")
}

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quite:
		return errors.New("workpool is quite")
	case p.tasks <- t:
		return nil
	}
}
func New(capacity int) *Pool {
	p := &Pool{
		capacity: capacity,
		active:   make(chan struct{}, capacity),
		tasks:    make(chan Task),
		quite:    make(chan struct{}),
	}
	fmt.Println("workerpool1 start ###########")
	go p.run()
	return p
}
