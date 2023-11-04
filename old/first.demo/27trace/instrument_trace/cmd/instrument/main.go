package main

import (
	"flag"
	"fmt"
	"hz.com/yk/instrument_trace/instrumenter"
	"hz.com/yk/instrument_trace/instrumenter/ast"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	wrote bool
)

func init() {
	flag.BoolVar(&wrote, "w", false, "write result to  (source) file insted of stdout")
}

func usage() {
	fmt.Printf("instrument [-w] xxx.go")
	flag.PrintDefaults()
}

func main() {
	fmt.Println(os.Args)
	flag.Usage = usage
	//使用标准库的 flag 包实现对命令行参数（这里是 -w）的解 析
	flag.Parse()

	if len(os.Args) < 2 {
		usage()
		return
	}
	var file string
	if len(os.Args) == 3 {
		file = os.Args[2]
	}

	if len(os.Args) == 2 {
		file = os.Args[1]
	}
	if filepath.Ext(file) != ".go" {
		usage()
		return
	}
	// 声明instrumenter.Instrumenter接口类型变量
	var ins instrumenter.Instrumenter
	// 创建以ast方式实现Instrumenter接口的ast.instrumenter实例
	ins = ast.New("hz.com/yk/instrument_trace", "trace", "Trace")
	newSrc, err := ins.Instrument(file) // 向Go源文件所有函数注入Trace函数
	if err != nil {
		panic(err)
	}

	if newSrc == nil {
		// add nothing to the source file. no change
		fmt.Printf("no trace added for %s\n", file)
		return
	}

	if !wrote {
		// 将生成的新代码内容输出到stdout上
		fmt.Println(string(newSrc))
		return
	}

	// write to the source file
	// 将生成的新代码内容写回原Go源文件
	if err = ioutil.WriteFile(file, newSrc, 0666); err != nil {
		fmt.Printf("write %s error: %v\n", file, err)
		return
	}
	fmt.Printf("instrument trace for %s ok\n", file)
}
