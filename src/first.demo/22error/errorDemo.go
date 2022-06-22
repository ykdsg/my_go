package main

import (
	"errors"
	"fmt"
)

var ErrSentinel = errors.New("the underlying sentinel error")

func errorIs() {
	error1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
	error2 := fmt.Errorf("wrap err1: %w", error1)
	println("error2 == ErrSentinel:", error2 == ErrSentinel)
	if errors.Is(error2, ErrSentinel) {
		println("err2 is ErrSentinel")
		return
	}
	println("err2 is not ErrSentinel")
}

type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

func errorAs() {
	var err = &MyError{"MyError error demo"}
	err1 := fmt.Errorf("wrap err: %w", err)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	var e *MyError
	//通过As函数给错误处理方检视错误值。
	if errors.As(err2, &e) {
		println("MyError is on the chain of err2")
		println("e == err:", e == err)
		return
	}
	println("MyError is not on the chain of err2")
}

func main() {
	errorIs()
	println("1111111111111111111111111........................")
	errorAs()
}
