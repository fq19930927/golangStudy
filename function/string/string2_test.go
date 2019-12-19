package string

import (
	"strconv"
	"strings"
	"testing"
)

//字符串分割
func TestStringSplit(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Logf(part)
	}
	//拼接
	t.Log(strings.Join(parts, "-"))
}

//转换测试
func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("s = " + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(i + 10)
	}
}
