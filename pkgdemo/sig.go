package pkgdemo

import (
	"os"
	"os/signal"
	"syscall"
)

func sigDemo() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
	// 执行程序收尾工作。。。
}
