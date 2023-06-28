package question

import (
	"fmt"
	"math/rand"
)

// 在 golang 协程和channel配合使用
// 写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

func printRands() {
	dataCh, done := make(chan int, 5), make(chan struct{})
	go func() {
		for i := 0; i < 5; i++ {
			dataCh <- rand.Intn(10)
		}
		close(dataCh)
	}()

	go func() {
		for value := range dataCh {
			fmt.Println(value)
		}
		close(done)
	}()
	<-done
}
