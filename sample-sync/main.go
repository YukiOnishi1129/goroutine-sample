package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello, from %v!\n", id)
}

func main() {
	var wg sync.WaitGroup

	// 1と2が並行で処理される
	//Addの呼び出しは、監視対象のgoroutineの関数の外で行うようにすること
	// goroutineが起動する前に、Waitが先に呼び出されて処理が終了される可能性があるため
	//wg.Add(1)
	//go func() {
	//	//関数の終わりにDoneする
	//	defer wg.Done()
	//	fmt.Println("1st goroutine start")
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("1st goroutine end")
	//}()
	//wg.Add(1)
	//go func() {
	//	//関数の終わりにDoneする
	//	defer wg.Done()
	//	fmt.Println("2st goroutine start")
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("2st goroutine end")
	//}()
	//
	//wg.Wait()

	//
	var CPU int = runtime.NumCPU()

	// 呼び出し回数がわかっているなら、forループの外で呼び出しても良い
	wg.Add(CPU)

	for i := 1; i <= CPU; i++ {
		//forループでgoroutineを呼び出すごとにAddするのが慣例だけど
		//wg.Add(1)
		go Hello(&wg, i)
	}

	wg.Wait()

}
