package offer

import "math"

var min float64 = 0.0000001
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if equals(x, 0) {
		return 0
	}
	b := true
	if n < 0 {
		n = 0-n
		b =false
	}
	var res float64 = 0
	if n&1 ==1{
		pow := myPow(x, (n-1)/2)
		res = pow* pow *x
	}else {
		pow := myPow(x, n/2)
		res = pow * pow
	}
	if b {
		return res
	} else {
		return 1/res
	}
}

func equals(x, y float64) bool {
	if math.Abs(x-y) < min {
		return true
	}
	return false
}
