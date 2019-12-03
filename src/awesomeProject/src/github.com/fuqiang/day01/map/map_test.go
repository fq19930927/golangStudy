package _map

import "testing"

//map初始化
func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1[1])
	t.Logf("len1 = %d", len(m1))
	m2 := map[int]int{}
	m2[3] = 3
	t.Log(m2[999])
	t.Logf("len2 = %d", len(m2))
	m3 := make(map[int]int, 10)
	m3[1] = 1
	t.Logf("len3 = %d", len(m3))
}

//空值test
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	m1[3] = 0
	if _, v := m1[3]; v {
		t.Log("key 3 exists")
	} else {
		t.Log("key 3 does not exist")
	}
}

//遍历测试
func TestRangeTest(t *testing.T) {
	m := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m {
		t.Log(k, v)
	}
}

//测试map的value为方法
func TestValueFunc(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}
