package permute

import (
	"fmt"
	"testing"
)

func TestP(t *testing.T) {
	l := permute([]int{3,2,1})
	fmt.Println(l)
}
