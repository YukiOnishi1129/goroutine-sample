package main

import (
	"fmt"
	"sync"
)

func main() {
	//順番を制御する際はWaitGroupを使
	//WaitGroupは、ゴルーチンの完了を待つための仕組み
	var wg sync.WaitGroup
	//競合状態を防ぐために排他処理を実施する
	//排他処理を実施するためには、ミューテックスを使う
	//ミューテックスは、排他処理を実施するための仕組み
	var memoryAccess sync.Mutex
	var data int

	wg.Add(1)

	go func() {
		defer wg.Done()
		//ロックとアンロックを多用すると処理が重くなるので使い所は注意
		//ロックしてからデータにアクセス
		memoryAccess.Lock()
		//クリティカルセクション
		data++
		//アクセスしたらアンロックする
		memoryAccess.Unlock()
	}()

	//Addされたゴルーチンを確実に待つ
	wg.Wait()

	//根本的には解決しない方法
	//time.Sleep(1 + time.Second)

	memoryAccess.Lock()
	if data == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(data)
	}
	memoryAccess.Unlock()
}

//競合状態
// 複数のゴルーチンがあって、実行の順番が定まっていない時に実行して予期せぬ値になること

//競合状態
//複数のゴルーチンが同じ変数にアクセスするときに、
//その変数の値が予期せぬ値になることを競合状態という。
//競合状態を避けるためには、複数のゴルーチンが同じ変数にアクセスするときに、
//その変数の値を変更するときは排他制御を行う必要がある。

//排他制御を行う方法は、ミューテックスやチャネルを使う方法がある。
//ミューテックスは、複数のゴルーチンが同じ変数にアクセスするときに、
//その変数の値を変更するときは排他制御を行う必要がある。
