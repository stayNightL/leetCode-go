package longestCommonPrefix

import (
	"fmt"
	"testing"
)

func TestLongest(t *testing.T) {
	s := [3]string{"dog", "dacecar", "dar"}
	fmt.Println(longestCommonPrefix(s[:]))
}
func TestLongest1(t *testing.T) {
	s := [1]string{"f"}
	fmt.Println(longestCommonPrefix(s[:]))
}
