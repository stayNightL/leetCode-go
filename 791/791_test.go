package customSortString

import (
	"fmt"
	"testing"
)

func TestCustomSortString(t *testing.T) {
	var s = "abc"
	var ts = "dcba"
	var str = customSortString(s, ts)
	fmt.Printf(str)
}
