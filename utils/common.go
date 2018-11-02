package utils

import (
	"github.com/cocopc/gcommons/log"
	"net"
	"os"
	"os/signal"
	"syscall"
)


// 等待退出信号
func WaitExitSign(){
	c :=make(chan os.Signal,1)
	// Interrupt SiGINT 都是 Ctrl-C
	signal.Notify(c,os.Interrupt,os.Kill,syscall.SIGINT,syscall.SIGKILL,syscall.SIGTERM,syscall.SIGHUP)
	l.Log("GET ExitSign !")
	<-c
}


func GetIp4Byname(host string) (ips []string, err error) {
	addrs, err := net.LookupIP(host)
	if err != nil {
		return
	}
	ips = make([]string, len(addrs))
	for i, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ips[i] = ipv4.String()
		}
	}
	return
}
