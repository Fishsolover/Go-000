package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var CloseChan chan bool

//开启一个HTTP服务
func startHttpServer() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("响应请求成功！")
	})

	http.HandleFunc("/close", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("开始关闭端口")
		CloseChan <- true
	})

	err := http.ListenAndServe(":9112", nil)
	return err
}

//关闭一个HTTP服务
func closeHttpServer() error {
	select {
	case <-CloseChan:
		return errors.New("退出")
	}
}

type Signal struct {
	SignalChan chan os.Signal
	SignalDone chan bool
	Actions    []SignalAction
}

type SignalAction interface {
	Name() string
	Action(os.Signal) error
}

//监听
func listenSignals() error {
	var signalMain Signal
	signalMain.SignalChan = make(chan os.Signal, 1)
	signalMain.SignalDone = make(chan bool, 1)
	signalMain.Actions = registerActions()
	signal.Notify(signalMain.SignalChan, syscall.SIGINT)
	var sg os.Signal
	select {
	case sg = <-signalMain.SignalChan:
		return errors.New(sg.String())
	}
	return nil
}

//注册动作

func registerActions() []SignalAction {
	var signalActions []SignalAction
	//此处开启注册
	return signalActions
}

func main() {

	CloseChan = make(chan bool, 1)

	group, _ := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return startHttpServer()
	})

	group.Go(func() error {
		return closeHttpServer()
	})

	group.Go(func() error {
		return listenSignals()
	})

	if err := group.Wait(); err != nil {
		panic(err.Error())
	} else {
		panic("")
	}

}
