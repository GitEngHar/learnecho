package main

import (
	"fmt"
	"time"
)

// closeされたchannelにデータを送ることはできない
func add_value_to_close() {
	ch := make(chan int)
	close(ch)
	ch <- 1 // send on closed channel error
}

// closeしておらず永久に待ち続ける
func reak_gorutine() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()
	for v := range ch {
		fmt.Println(v) // fatal error: all goroutines are asleep - deadlock!
	}

	time.Sleep(time.Second)
}

// リクエストがキャパを超えている
func many_request() {
	ch := make(chan int, 3)
	// リクエスト側
	go func() {
		// too many request
		// 読み取り側 3つ目以降は受信されない限りブロック
		for i := 0; i < 10; i++ {
			ch <- i
		}
		defer close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
func main() {
	many_request()
}
