package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

var CloseChan chan bool

//开启一个HTTP服务
func startHttpServer() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("响应请求成功！")
	})

	http.HandleFunc("/close", func(writer http.ResponseWriter, request *http.Request) {
		CloseChan <- true
	})

	err := http.ListenAndServe(":9112", nil)
	return err
}

//关闭一个HTTP服务
func closeHttpServer() error {
	if <-CloseChan {
		return errors.New("退出")
	}
	return nil
}

//监听

func main() {

	group, _ := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return startHttpServer()
	})

	group.Go(func() error {
		return closeHttpServer()
	})

	if err := group.Wait(); err != nil {
		panic(err.Error())
	}

}
