package test

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	c, d := sum(1, 2)
	fmt.Println(c, d)
}
func sum(a int, b int) (c int, d int) {
	c = 1
	d = 2
	return 0, 1
	return a + b, a - b
}
