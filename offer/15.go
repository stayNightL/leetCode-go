package offer


func hammingWeight(num uint32) int {
	sum := 0
	for  {
		if num == 0 {
			return sum
		}
		num = num&(num-1)
		sum ++
	}
}
