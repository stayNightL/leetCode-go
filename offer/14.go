package offer

func cuttingRope(n int) int {
	if n == 0 || n==1  {
		return n
	}
	ints := make([]int, n)
	ints[0] =0
	ints[1] =1
	max := 0
	for i := 1; i < n; i++ {
		i2:= i*cuttingRopeN(ints,n-i)
		i3 := i * (n - i)
		if max < i3 {
			max = i3
		}
		if max < i2 {
			max = i2
		}
	}
	return max
}

func cuttingRopeN(ints []int, n int) int {
	if ints[n] != 0 {
		return  ints[n]
	}
	max := 0
	for i := 1; i < n; i++ {
		i2:= i*cuttingRopeN(ints,n-i)
		i3 := i * (n - i)
		if max < i3 {
			max = i3
		}
		if max < i2 {
			max = i2
		}
	}
	ints[n] = max
	return max
}

