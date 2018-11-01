package utils

import (
	"os"
	"os/signal"
	"syscall"
)


// 等待退出信号
func WaitExitSign(){
	c :=make(chan os.Signal,1)
	// Interrupt SiGINT 都是 Ctrl-C
	signal.Notify(c,os.Interrupt,os.Kill,syscall.SIGINT,syscall.SIGKILL,syscall.SIGTERM,syscall.SIGHUP)
	<-c
}