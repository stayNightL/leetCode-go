package permute

import (
	"fmt"
	"testing"
)

func TestP(t *testing.T) {
	l := permuteUnique([]int{1, 2, 1})
	fmt.Println(l)
}
