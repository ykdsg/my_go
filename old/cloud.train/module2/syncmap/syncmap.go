package main

import (
	"sync"
	"time"
)

func main() {
	//unsafeWrite()
	safeWrite()
	time.Sleep(time.Second)
}

type SafeMap struct {
	safeMap map[int]int
	sync.Mutex
}

func (s SafeMap) Write(k int, v int) {
	s.Lock()
	defer s.Unlock()
	s.safeMap[k] = v
}

func safeWrite() {
	s := SafeMap{
		safeMap: map[int]int{},
		Mutex:   sync.Mutex{},
	}
	for i := 0; i < 100; i++ {
		go func() {
			s.Write(1, 1)
		}()
	}
}

func unsafeWrite() {
	conflictMap := make(map[int]int)
	for i := 0; i < 100; i++ {
		go func() {
			conflictMap[1] = i
		}()
	}
}
