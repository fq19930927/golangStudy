package operator_test

import "testing"

//数组比较
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 3, 2, 4}
	t.Log(a == b, b == c)
}
