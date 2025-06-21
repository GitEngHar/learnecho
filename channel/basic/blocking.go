package main

import (
	"fmt"
	"time"
)

func main() {
	// cap 2
	ch := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("send to channel")
			ch <- i
		}
		close(ch) //これがないとチャネルからの値を待ち続ける。永遠にブロッキングとなる。
	}()

	time.Sleep(1 * time.Second)
	for msg := range ch {
		fmt.Println("Received %d", msg)
	}
}
