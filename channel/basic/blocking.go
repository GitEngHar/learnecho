package main

import (
	"fmt"
	"time"
)

func main() {
	// cap 2
	ch := make(chan int, 2) // capが送信されるデータよりも少ないため、空きができるまで待ちとなる
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("send to channel")
			ch <- i
		}
		defer close(ch) //これがないとチャネルからの値を待ち続ける。永遠にブロッキングとなる。
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Sleep 1s")
	for msg := range ch { //チャネルが閉じるまでデータ取得する
		fmt.Println("Received %d", msg)
	}

}
