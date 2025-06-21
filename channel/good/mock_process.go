package main

import (
	"fmt"
	"sync"
	"time"
)

const numberOfProcess = 2

func main() {
	requests := make(chan int, 5)
	var wg sync.WaitGroup
	for i := 0; i < numberOfProcess; i++ {
		wg.Add(1)
		// チャネルからデータの読み取り
		go worker(requests, &wg)
	}
	// チャネルへのデータリクエスト
	request(requests)
	wg.Wait()
	fmt.Println("Process all done")
}

func request(requests chan int) {
	// 0~9までの値を全てchannelに送信する
	for i := 0; i < 10; i++ {
		requests <- i
	}
	defer close(requests)
}

func worker(requests chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for req := range requests {
		process(req) // channelから値を取得する
	}
}

func process(req int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Processed request: %d\n", req)
}
