package constant_test

import "testing"

const (
	Monday = iota + 1;
	Tuesday
	Wenesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCons(t *testing.T) {
	t.Log(Monday, Tuesday)
	t.Log(Readable, Writable, Executable)
}
