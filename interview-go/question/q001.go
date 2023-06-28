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
	numberCh, charCh, closeCh := make(chan struct{}), make(chan struct{}), make(chan struct{})

	go func() {
		current := 1
		for {
			select {
			case <-numberCh:

				fmt.Print(current)
				current++
				fmt.Print(current)
				current++

				charCh <- struct{}{}
			case <-closeCh:
				return
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		current := 'A'
		for range charCh {
			if current > 'Z' {
				wg.Done()
				close(closeCh)
				return
			}

			fmt.Print(string(current))
			current++
			fmt.Print(string(current))
			current++

			numberCh <- struct{}{}
		}
	}()

	numberCh <- struct{}{}
	wg.Wait()
	close(numberCh)
	close(charCh)
}
