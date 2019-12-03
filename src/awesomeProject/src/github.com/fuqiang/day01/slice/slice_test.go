package slice

import "testing"

// slice是可变长度的
func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0 ,1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1,2,3,4}
	t.Log(len(s1), cap(s1))

	s2:=make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
}

func TestSet(t *testing.T)  {
	mySet := map[int]bool{}
	n :=1
	mySet[n] = true
	if mySet[1] {
		t.Log("key 1 exists")
	} else {
		t.Log("key 1 does not exist")
	}
	t.Log(len(mySet))
	delete(mySet,1)
	t.Log(len(mySet))
}
