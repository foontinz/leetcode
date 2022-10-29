package solutions

import "fmt"

func SolutionSparse(N int) int {
	for firstComp := 0; firstComp <= N/2; firstComp++ {
		secondComp := N - firstComp

		firstCompBin := fmt.Sprintf("%b", firstComp)
		secondCompBin := fmt.Sprintf("%b", secondComp)

		firstFlag := true
		secondFlag := true

		lastBit := '0'
		for _, bit := range firstCompBin {
			if bit == '1' && lastBit == '1' {
				firstFlag = false
				break
			}
			lastBit = bit
		}

		lastBit = '0'
		for _, bit := range secondCompBin {
			if bit == '1' && lastBit == '1' {
				secondFlag = false
				break
			}
			lastBit = bit
		}
		if firstFlag && secondFlag {
			return firstComp
		}
	}
	return -1
}
