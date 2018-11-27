package threeSum

import (
	"fmt"
	"testing"
)

func TestTreeNode(t *testing.T) {
	numbs := [4]int{1, 2, 3, 4}
	

	fmt.Printf(buildTree(numbs[:]).toString())

}
