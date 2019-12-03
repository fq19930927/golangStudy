package arr

import "testing"

// 数组操作
func TestArr(t *testing.T) {
	var arr [3] int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{5, 6, 7, 8}
	t.Log(arr, arr1, arr2)
	arr1[1] = 5
	t.Log(arr1)
	for _, e := range arr2 {
		t.Log(e)
	}
	t.Log(arr2[1:3])
	t.Log(arr2[1:len(arr2)])
	t.Log(arr2[1:])
	t.Log(arr2[:3])
}

// 数组切片
func TestArrSection(t *testing.T)  {
	arr := [...]int{1,2,3,4, 5}
	arr2 := arr[:3]
	t.Log(arr2)
}
