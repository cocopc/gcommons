package app

import (
	"github.com/cocopc/gcommons/log"
	"github.com/cocopc/gcommons/utils"
	"os"
)


func init(){
	l.Define("[gcommons]",l.Blue,"INFO")
}

// Run 运行一个应用，需要初始化函数，运行作业函数，清理资源函数，收到退出信号时清理资源
func Run(appName string,initFunc,jobFunc,cleanupFunc func() error)  {
	l.Printf("Initial [%s]", appName)
	if err := initFunc(); err != nil {
		l.Printf("Initial [%s] failure: [%s]", appName, err)
		panic(err)
	}
	l.Printf("Initial [%s] complete", appName)
	go func() {
		if err := jobFunc(); err != nil {
			l.Printf("[%s] run error: [%v]", appName, err)
			panic(err)
		}
	}()

	utils.WaitExitSign()
	l.Printf("[%s] watched the exit signal, start to clean", appName)
	if err := cleanupFunc(); err != nil {
		l.Printf("[%s] clean failed: [%v]", appName, err)
		panic(err)
	}
	l.Printf("[%s] clean complete, exited", appName)

	os.Exit(0)

}