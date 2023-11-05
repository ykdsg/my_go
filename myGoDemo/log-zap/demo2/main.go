package main

import "github.com/yk/zap-usage/pkg/log"

func main() {
	log.Info("demo1:", log.String("app", "start ok"))
}
