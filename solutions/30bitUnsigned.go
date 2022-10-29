package solutions

import (
	"math/bits"
)

func Solution30bitUnsigned(A int, B int, C int) int {
	possibilitiesA := 1 << (30 - bits.OnesCount(uint(A)))
	possibilitiesB := 1 << (30 - bits.OnesCount(uint(B)))
	possibilitiesC := 1 << (30 - bits.OnesCount(uint(C)))

	possibilitiesAorB := 1 << (30 - bits.OnesCount(uint(A|B)))
	possibilitiesAorC := 1 << (30 - bits.OnesCount(uint(A|C)))
	possibilitiesBorC := 1 << (30 - bits.OnesCount(uint(B|C)))
	possibilitiesAorBorC := 1 << (30 - bits.OnesCount(uint(A|B|C)))

	return possibilitiesA + possibilitiesB + possibilitiesC - possibilitiesAorB - possibilitiesAorC - possibilitiesBorC - possibilitiesAorBorC

}
