package validPalindrome

import (
	"fmt"
	"testing"
)

func TestIsvalid(t *testing.T) {
	var s = "aba"
	if isValid(s) {
		fmt.Print("yes,itis")
	}
}
func TestValidPalindrome(t *testing.T) {
	var s = "yd"
	if validPalindrome(s) {
		fmt.Print("yes,itis")
	}
}
//TEST
func TestSlice(t *testing.T) {
	var s = "abc"
	for i, r := range s {
		fmt.Printf("r=" +string(r) + "===" + s[:i] + s[i+1:])
	}
}