package threeSum

import (
	"container/list"
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
	for v := sli.Front(); v != nil; v.Next() {

		fmt.Println(v.Value)
	}

}

func TestSlice(t *testing.T) {
	s := make([]int, 10)
	s = append(s, 1)
	printSlice(s)
}

type e struct {
	Value int
}

func TestList(t *testing.T) {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushBack(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
