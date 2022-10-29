package solutions

import (
	"bytes"
	"fmt"
)

func ConvertZigZag(s string, numRows int) {
	var wordMatrix [][]uint8

	for i := 0; i < numRows; i++ {
		wordMatrix = append(wordMatrix, []uint8{})
		for j := 0; j < len(s); j++ {
			wordMatrix[i] = append(wordMatrix[i], 0)
		}
	}

	numRows--
	char := 0

	for i, j := 0, 0; char < len(s); {
		if i%numRows != 0 || i == 0 {
			wordMatrix[i][j] = s[char]
			char++
			i++
		} else {
			for i != 0 && char < len(s) {
				wordMatrix[i][j] = s[char]
				char++
				i--
				j++

			}
		}
	}

	var result string
	for _, charArr := range wordMatrix {
		for _, c := range charArr {
			if c != 0 {
				result += string(c)
			}
		}
	}
	println(result)
}

func ConvertZigZag2(s string, numRows int) {
	if numRows == 1 {
		fmt.Println(s)
	}
	var b bytes.Buffer
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); {
			b.WriteByte(s[j])
			j += numRows + numRows - 2
			if i == 0 || i == numRows-1 {
				continue
			}
			j -= i + i
			if j >= len(s) {
				break
			}
			b.WriteByte(s[j])
			// Set j back again to the value in the same row and the next column
			j += i + i
		}
	}
	fmt.Println(b.String())

}
