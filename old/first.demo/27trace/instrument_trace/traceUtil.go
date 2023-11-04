package trace

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

var goroutineSpace = []byte("goroutine ")

//获取Goroutine ID，类似线程id
func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, goroutineSpace)

	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}

//func main() {
//	defer Trace()()
//	foo()
//}
