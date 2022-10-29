package solutions

import (
	"fmt"
)

func SolutionBinaryGap(N int) int {
	binaryRepr := fmt.Sprintf("%b", N)
	gaps := []int{}
	flag := false
	var counter int
	for _, char := range binaryRepr {
		if char == '1' {
			flag = true
			gaps = append(gaps, counter)
			counter = 0
			continue
		}
		if flag {
			counter++
		}
	}
	res := 0
	for _, gap := range gaps {
		if gap > res {
			res = gap
		}
	}
	return res
}
