package pkg1

import "github.com/yk/zap-usage/pkg/log"

func Foo() {
	log.Info("call foo", log.String("url", "https://tonybai.com"),
		log.Int("attempt", 3))
}
