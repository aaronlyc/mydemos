package question

import (
	"fmt"
	"sync"
)

// 交替打印数字和字母
// 问题描述
// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

func altPrint() {
	letter, number := make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for range number {
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letter <- struct{}{}
		}
	}()

	wg.Add(1)
	go func() {
		i := 'A'
		for range letter {
			if i > 'Z' {
				wg.Done()
				return
			}
			fmt.Printf("%c", i)
			i++
			fmt.Printf("%c", i)
			i++
			number <- struct{}{}
		}
	}()

	number <- struct{}{}
	wg.Wait()
}
