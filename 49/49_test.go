package groupanagrams

import (
	"fmt"
	"testing"
)

func TestGru(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
func TestGru2(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"c", "c"}))
}
