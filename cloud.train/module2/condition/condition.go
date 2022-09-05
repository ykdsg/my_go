package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []string
	cond  *sync.Cond
}

func (q *Queue) Enqueue(str string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, str)
	fmt.Printf("putting %s to queue,notify all \n", str)
	//相当于java的notifyAll
	q.cond.Broadcast()
	//相当于java 的notify
	//q.cond.Signal()
}

func (q *Queue) Dequque() (str string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 {
		fmt.Println("no data available,wait ")
		q.cond.Wait()
	}
	str, q.queue = q.queue[0], q.queue[1:]
	return
}

func main() {
	q := Queue{
		queue: make([]string, 0),
		cond:  sync.NewCond(&sync.Mutex{}),
	}
	go func() {
		for {
			q.Enqueue("a")
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		q.Dequque()
		time.Sleep(time.Second)
	}
}
