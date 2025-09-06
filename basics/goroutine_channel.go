package main

import (
	"fmt"
	"time"
)

func worker(name string, ch chan string) {
	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("%s: %d", name, i)
		ch <- msg // channel に送信
		time.Sleep(400 * time.Millisecond)
	}
	close(ch) // 終わったらクローズ
}

func main() {
	ch := make(chan string)

	go worker("Worker", ch)

	// channel から受け取る
	for msg := range ch {
		fmt.Println("受信:", msg)
	}
}

