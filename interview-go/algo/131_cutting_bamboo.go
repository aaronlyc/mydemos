package algo

import "math"

func cuttingBamboo(bamboo_len int) int {
	if bamboo_len <= 3 {
		return bamboo_len - 1
	}

	q := bamboo_len / 3
	r := bamboo_len % 3

	if r == 0 {
		return int(math.Pow(3, float64(q)))
	} else if r == 1 {
		return int(math.Pow(3, float64(q-1)) * 4)
	} else {
		return int(math.Pow(3, float64(q)) * 2)
	}
}
