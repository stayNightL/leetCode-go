package findDiagonalOrder

import (
	"testing"
)

type s1 []int
type s2 []s1

func TestSize(t *testing.T) {

	s := [][]int{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{7, 8, 9, 10, 11, 12},
		[]int{13, 14, 15, 16, 17, 18},
		[]int{20, 21, 22, 23, 24, 25}}
	findDiagonalOrder(s[:][:])
}
