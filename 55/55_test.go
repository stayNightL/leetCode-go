package canJump

import (
	"fmt"
	"testing"
)

func TestCanJ(t *testing.T) {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}
func TestCanJ2(t *testing.T) {
	fmt.Println(canJump([]int{3,2,1,0,4}))
}