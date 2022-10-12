package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

//golog 默认不是打印到标准输出，需要增加2个参数
//-logtostderr：打印到标准错误而不是文件。
//-alsologtostderr：同时打印到标准错误。
func main() {
	flag.Set("v", "4")
	flag.Parse()
	//defer glog.Flush()
	//glog.V(2).Info("Starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}
