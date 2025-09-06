package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(name string, wg *sync.WaitGroup) {
	defer wg.Done() // 終わったらカウントを減らす

	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // 2つのgoroutineを待つ
	go worker("A", &wg)
	go worker("B", &wg)

	// メイン処理
	worker("Main", &wg) // これは待たなくてもいいなら wg なしでもOK

	wg.Wait() // 全部終わるまで待機
}

