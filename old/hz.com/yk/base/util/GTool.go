package util

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

//因为在断点里面的log调用的时候，如果工程中没有调用过这个函数将会报错，所以这里先调用一下。
func init() {
	GoID()
}

func GoID() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return strconv.Itoa(id)
}
