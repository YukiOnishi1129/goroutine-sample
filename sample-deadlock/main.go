package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
	name  string
}

// デッドロックを起こす処理
func main() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		fmt.Printf("%v がロックを取得しました\n", v1.name)
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)

		v2.mu.Lock()
		fmt.Printf("%v がロックを取得しました\n", v2.name)
		defer v2.mu.Unlock()

		fmt.Println(v1.value + v2.value)
	}

	var a value = value{name: "a"}
	var b value = value{name: "b"}

	wg.Add(2)

	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}

//デッドロック
//複数の並行処理が完了するのを待っていて、止まってしまうこと

//デットロック
//デットロックとは、複数のゴルーチンが相互に待ち合ってしまい、
//全てのゴルーチンがブロックされてしまう状態のことです。
//デットロックが発生すると、プログラムは終了せずに永遠にブロックされ続けます。
//デッドロックの例として、以下のようなゴルーチンがあるとします。
//このゴルーチンは、チャネルに値を送信する処理を無限ループで繰り返しています。
//このゴルーチンを起動すると、チャネルに値を送信する処理が実行されます。
//しかし、チャネルに値を送信する処理は、チャネルに値を受信する処理が実行されるまでブロックされます。
//そのため、このゴルーチンは、チャネルに値を送信する処理を実行した後、
//チャネルから値を受信する処理が実行されるまでブロックされます。
//このように、チャネルに値を送信する処理とチャネルから値を受信する処理が
//相互にブロックしあってしまう状態のことをデッドロックと呼びます。
//デッドロックが発生すると、プログラムは終了せずに永遠にブロックされ続けます。
//デッドロックを回避するには、チャネルに値を送信する処理とチャネルから値を受信する処理の
//どちらか一方のみが実行されるようにする必要があります。
