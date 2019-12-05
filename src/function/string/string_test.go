package string

import "testing"

func TestString(t *testing.T) {
	var s string
	s = "hello"
	t.Log(len(s))
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	//len求的是string中的byte长度
	//string是不可变的byte slice
	t.Log(len(s))

	c := []rune(s)
	t.Log(len(c))
	t.Logf("unicode, %x", c)
	t.Logf("utf-8, %x", s)

}

func TestStringToRune(t *testing.T) {
	s := "我爱我的祖国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
