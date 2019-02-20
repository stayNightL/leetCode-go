package divide

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	fmt.Println(divide(10, 2))
}
func TestDivide11(t *testing.T) {
	fmt.Println(divide(1, -1))
}

func TestDivide12(t *testing.T) {
	fmt.Println(divide(100, 2))
}
func TestDivide1222222(t *testing.T) {
	fmt.Println(divide(2147483647, 1))
}
