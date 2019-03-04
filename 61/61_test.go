package rotateRight

import (
	"fmt"
	"testing"
)

func TestRotateRight(t *testing.T) {
	fmt.Println(rotateRight(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 1))
}
