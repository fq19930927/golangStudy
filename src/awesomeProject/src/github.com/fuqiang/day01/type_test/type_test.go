package type_test

import "testing"

type MyInt = int64;

//隐式类型转换测试
func TestImplicit(t *testing.T) {
	var a int = 1;
	var b MyInt;
	b = MyInt(a)
	t.Log(b)
}

//指针测试
func TestPoint(t *testing.T) {
	a := 1
	b := &a
	t.Log(a, b)
	t.Logf("%T %T", a, b)
}

//测试String
func TestString(t *testing.T)  {
	var s string//go默认String类型是""
	t.Log("*"+s + "*")
	t.Log(len(s))
}
