package func_test

import (
	"math/rand"
	"testing"
)

func returnMultiVal() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, _ := returnMultiVal()
	t.Log(a)
}
