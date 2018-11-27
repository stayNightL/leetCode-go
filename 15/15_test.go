package threeSum

import (
	"fmt"
	"testing"
)

func TestTreeNode(t *testing.T) {
	numbs := [4]int{1, -2, 3, 4}

	fmt.Printf(buildTree(numbs[:]).toString())

}

func TestEasy(t *testing.T) {
	numbs := [4]int{1, -2, 3, 4}
	s := buildTree(numbs[:])
	sli := s.easytravell()
	for _, v := range sli {
		fmt.Println(v.nodeString())
	}

}

func TestSlice(t *testing.T) {
	s := make([]int, 10)
	s = append(s, 1)
	printSlice(s)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
