package test

import (
	"log"
	"runtime"
	"testing"
	"time"
)

// 只有一个sender的情况，直接从sender关闭即可
func Test1(t *testing.T) {
	dataCh := make(chan int, 100)

	// sender
	go func() {
		for i := 0; i < 1000; i++ {
			dataCh <- i + 1
		}
		log.Println("send complete")
		close(dataCh)
	}()

	//receiver
	for i := 0; i < 5; i++ {
		go func() {
			for {
				data, ok := <-dataCh
				if !ok { // 已关闭
					return
				}
				_ = data
			}
		}()
	}

	select {
	case <-time.After(time.Second * 5):
		log.Println("runtime: ", runtime.NumGoroutine())
	}
}

func Test2(t *testing.T) {
	dataCh := make(chan int, 100)

	// sender
	go func() {
		for i := 0; i < 1000; i++ {
			dataCh <- i + 1
		}
		log.Println("send complete")
		close(dataCh)
	}()

	//receiver
	for i := 0; i < 5; i++ {
		go func() {
			for {
				data, ok := <-dataCh
				if !ok { // 已关闭
					return
				}
				_ = data
			}
		}()
	}

	select {
	case <-time.After(time.Second * 5):
		log.Println("runtime: ", runtime.NumGoroutine())
	}
}
