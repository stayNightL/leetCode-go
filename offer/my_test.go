package offer

import (
	"fmt"
	"testing"
)

func TestT11(t *testing.T) {
	i := []int{3,1,3,3,3}
	minArray(i)
}

func minArray(numbers []int) int {
	fmt.Println(numbers)
	l := 0
	if l =len(numbers);l == 0 {
		return 0
	}
	if numbers[0] < numbers[l-1] {
		return numbers[0]
	}
	//如果0 《 last，返回0
	if l == 2  {
		if numbers[0] < numbers[1] {
			return numbers[0]
		}else {
			return numbers[1]
		}
	}
	//二分法，记
	index := 0
	lastIndex := l-1
	midIndex := lastIndex/2
	if midIndex == 0 {
		return numbers[0]
	}
	fmt.Println(index,lastIndex,midIndex)
	if numbers[index] < numbers[midIndex]{
		return minArray(numbers[midIndex:])
	}else if numbers[midIndex] < numbers[lastIndex] {

		return minArray(numbers[:midIndex+1])
	}else if numbers[index] == numbers[midIndex]&&numbers[lastIndex] == numbers[midIndex]{
		return minArray(numbers[:lastIndex])
	} else if numbers[index] == numbers[midIndex]{
		return minArray(numbers[midIndex:])
	} else if numbers[lastIndex] == numbers[midIndex]{
		return minArray(numbers[:midIndex+1])
	}
	return 0
}

func exist(board [][]byte, word string) bool {
	y := len(board)
	if y == 0 {
		return false
	}
	x := len(board[0])
	if x == 0 {
		return false
	}
	res := make([][]int, y)
	for i := range res {
		res[i] = make([]int,x)
	}
	wordIndex := 0
	for y1 := 0; y1 < y; y1++ {
		for x1 := 0; x1 < x; x1++ {
			if dfs(board,res,word,y,x,y1,x1,wordIndex){
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, res [][]int, word string, y int, x int, y1 int, x1 int, wordIndex int) bool {
	if x1 < 0 || y1 < 0 || x1 >= x || y1 >= y ||res[y1][x1]==1 || word[wordIndex] != board[y1][x1] {
		return false
	}
	if wordIndex== len(word)-1 {
		return true
	}
	res[y1][x1] = 1
	wordIndex++
	res1 := dfs(board,res,word,y,x,y1+1,x1,wordIndex)||
		dfs(board,res,word,y,x,y1,x1+1,wordIndex)||
		dfs(board,res,word,y,x,y1-1,x1,wordIndex)||
		dfs(board,res,word,y,x,y1,x1-1,wordIndex)
	res[y1][x1] = 0
	return res1
}