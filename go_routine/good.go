package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	worker()
}

func worker() {
	const workerCounts = 3
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	//forが終わった後にgo routineが始まるので、3つのワーカーが起動しない
	//for w := 0; w < 3; w++ {
	//	go func() {
	//		for j := range jobs {
	//			results <- j * 2
	//			fmt.Printf("Worker %d\n", w)
	//		}
	//	}()
	//}

	//
	for w := 0; w < 3; w++ {
		var workLabel = "is working on job"
		go func(id int, label string) {
			for j := range jobs {
				fmt.Printf("Worker %d %s %d\n", id, workLabel, j)
				time.Sleep(1 * time.Second) // 仕事してるフリ
				results <- j * 2            // 結果を送る
			}
		}(w, workLabel)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 0; i < 10; i++ {
		fmt.Println(<-results)
	}
}

func timeOut() {
	// multiAPI
	timeOutFlag := false
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan string, 1)
	go func() {
		if timeOutFlag {
			time.Sleep(3 * time.Second)
		}
		ch <- "done"
	}()

	select {
	case <-ctx.Done():
		fmt.Println("timeout")
	case res := <-ch:
		fmt.Printf("res: %s", res)
	}
}

func multiAPI() {
	var wg sync.WaitGroup
	results := make(chan int, 2)

	wg.Add(2)

	start := time.Now()
	go func() {
		defer wg.Done()
		results <- fetchFromAPI1()
	}()

	go func() {
		defer wg.Done()
		results <- fetchFromAPI2()
	}()
	end := time.Now()
	wg.Wait()
	close(results)
	fmt.Printf("処理時間: %d", end.Sub(start))
	for r := range results {
		fmt.Println(r)
	}
}

// 1日分のデータのサマリ想定 (ServiceA)
func fetchFromAPI1() int {
	time.Sleep(2 * time.Second)
	return 25
}

// 1日分のデータのサマリ想定 (ServiceB)
func fetchFromAPI2() int {
	time.Sleep(1 * time.Second)
	return 15
}
