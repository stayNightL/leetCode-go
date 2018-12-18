package lengthOfLongestSubstring

import (
	"fmt"
	"testing"
)

func TestLength(t *testing.T) {
	fmt.Printf("%d", lengthOfLongestSubstring("abcabcabc"))
}
func TestLength2(t *testing.T) {
	fmt.Printf("%d", lengthOfLongestSubstring("bbb"))
}
func TestLength3(t *testing.T) {
	fmt.Printf("%d", lengthOfLongestSubstring("aabaab!bb"))
}
