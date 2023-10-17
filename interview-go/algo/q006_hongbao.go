package algo

import (
	"fmt"
	"math/rand"
)

// 让你设计一个微信发红包的api，你会怎么设计，不能有人领到的红包里面没钱，红包数值精确到分。
func redPackage(count int, money float64) {
	// 传入的小数点为2位
	m := int(100 * money)
	for i := count; i > 0; i-- {
		current := randomMoney(i, m)
		fmt.Printf("the people: %d, hongbao is: %.2f¥\n", i, float64(current)/100)
		m -= current
	}
}

func randomMoney(count, money int) int {
	if count == 1 {
		// fmt.Printf("the hongbao is: %.2f¥\n", float64(money)*0.01)
		return money
	}

	min := 1
	max := money / count * 2
	current := rand.Intn(max)
	if current < min {
		current = min
	}
	return current
}
