package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

type capitalizedReader struct {
	r io.Reader
}

func CapReader(r io.Reader) io.Reader {
	return &capitalizedReader{r: r}
}

func (r *capitalizedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}

	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}
	return n, err
}

func main() {
	r := strings.NewReader("hello,world!\n")
	//将 CapReader 和 io.LimitReader 串在了一起形成一条调用链，这条调用链的 功能变为：截取输入数据的前四个字节并将其转换为大写字母。
	r1 := CapReader(io.LimitReader(r, 4))
	if _, err := io.Copy(os.Stderr, r1); err != nil {
		log.Fatalln(err)
	}
}
