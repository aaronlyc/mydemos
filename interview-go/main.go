package main

import "fmt"

func main() {
	test := make([]int, 0, 10)

	f := func(t []int) {
		for i := 0; i < 3; i++ {
			t = append(t, i)
		}
	}

	f(test)

	fmt.Println(test[0:3])
}
