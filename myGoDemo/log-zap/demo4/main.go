package main

import "github.com/yk/zap-usage/pkg/log"

func main() {

	for i := 0; i < 20000; i++ {
		log.Info("demo3:", log.String("app", "start ok"),
			log.Int("major version", 3))
		log.Error("demo3:", log.String("app", "crash"),
			log.Int("reason", -1))
	}

}
