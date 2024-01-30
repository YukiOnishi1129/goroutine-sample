package main

import (
	"fmt"
	"runtime"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Hello")
}

func newConsumed() uint64 {
	runtime.GC()

	var s runtime.MemStats

	runtime.ReadMemStats(&s)

	return s.Sys
}

func main() {
	// goroutine is a lightweight thread of execution
	// ゴルーチンは非同期で実行される　= メイン処理が終わってもゴルーチンは実行され続ける = メイン処理内でゴルーチンを完了させたい
	// ゴルーチンの完了を待つためには、WaitGroupを使う

	//WaitGroup: mechanism to wait for waiting for completion of goroutine.
	//Add: add number of goroutines.
	//Done: notification of completion of goroutine.
	//Wait:wait for completion of goroutine.

	//var wg sync.WaitGroup
	//
	//wg.Add(1)
	//
	//go sayHello(&wg)
	//
	//wg.Add(1)
	//
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("Hello")
	//}()

	//wg.Wait()

	// Wait 2 seconds to complete the main process.
	// time.Sleep(2 * time.Second)

	//var wg sync.WaitGroup
	//
	//say := "Hello World"
	//
	//wg.Add(1)
	//
	////子供のgoroutineと親のgoroutineは同じアドレス空間で一考される
	//go func() {
	//	defer wg.Done()
	//	// クロージャーの中で上書きできる
	//	say = "Good Bye"
	//}()
	//
	//wg.Wait()
	//
	//fmt.Println(say)

	//var wg sync.WaitGroup
	//
	//tasks := []string{"A", "B", "C"}
	//
	//for _, task := range tasks {
	//	wg.Add(1)
	//	// 回避するには、引数に渡すtaskをコピーする
	//	go func(task string) {
	//		defer wg.Done()
	//		//3回"C"が出力される
	//		//forの中で展開されるtaskは同じアドレスなので上書きされる
	//		fmt.Println(task)
	//	}(task)
	//}
	//
	//wg.Wait()

	//	メモリの大きさを測定
	//	goroutineは完了しないとメモリの解放はされない

	var ch <-chan interface{}
	var wg sync.WaitGroup

	noop := func() {
		wg.Done()
		<-ch
	}

	const numGoroutines = 1000000

	wg.Add(numGoroutines)

	before := newConsumed()

	for i := 0; i < numGoroutines; i++ {
		go noop()
	}

	wg.Wait()

	after := newConsumed()

	//2.5kbほど
	//OSのスレッドは1~2MBほど
	//非常に軽い
	fmt.Println("%.3fkb", float64(after-before)/numGoroutines/1000)
}
