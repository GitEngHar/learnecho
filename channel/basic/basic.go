package main

import (
	"fmt"
	"time"
)

func main() {
	messageChannel := make(chan string)
	var message = "hello"
	go func() {
		time.Sleep(2 * time.Second)
		messageChannel <- "hello channel from go routine"
	}()
	message = <-messageChannel
	fmt.Println(message)
}
