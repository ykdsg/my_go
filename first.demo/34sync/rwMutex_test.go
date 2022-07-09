package main

import (
	"sync"
	"testing"
)

var cs1 = 0 // 模拟临界区要保护的数据
var mu1 sync.Mutex
var cs2 = 0 // 模拟临界区要保护的数据
var mu2 sync.RWMutex

//通过命令行参数 -cpu 控制GOMAXPROCS
//go test -cpu=2,8,16 -bench='mu$'  .
//如果通过goland，需要在program argument 中设置"-test.cpu=2,8,16"
func BenchmarkWriteSyncByMutex_mu(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu1.Lock()
			cs1++
			mu1.Unlock()
		}
	})
}

func BenchmarkReadSyncByMutex_mu(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu1.Lock()
			_ = cs1
			mu1.Unlock()
		}
	})
}
func BenchmarkWriteSyncByRWMutex_mu(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu2.Lock()
			cs2++
			mu2.Unlock()
		}
	})
}

func BenchmarkReadSyncByRWMutex_mu(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu2.RLock()
			_ = cs2
			mu2.RUnlock()
		}
	})
}
