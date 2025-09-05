package main

import (
	"fmt"
	"time"
)

// 仕事をする関数
func worker(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(500 * time.Millisecond) // 0.5秒待つ
	}
}

func main() {
	// goroutine で並列実行
	go worker("A")
	go worker("B")

	// メインも同時に動く
	worker("Main")
}

