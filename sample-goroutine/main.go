package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Hello")
}

func main() {
	// ゴルーチンは非同期で実行される　= メイン処理が終わってもゴルーチンは実行され続ける = メイン処理内でゴルーチンを完了させたい
	// ゴルーチンの完了を待つためには、WaitGroupを使う

	//WaitGroup: ゴルーチンの完了を待つための仕組み
	//Add: ゴルーチンの数を追加
	//Done: ゴルーチンの完了を通知
	//Wait: ゴルーチンの完了を待つ

	var wg sync.WaitGroup

	wg.Add(1)

	go sayHello(&wg)

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Hello")
	}()

	wg.Wait()

	//2秒間待ってmain処理を終了する
	//time.Sleep(2 * time.Second)
}
