package object

import (
	"fmt"
	"testing"
	"unsafe"
)

func main() {

}

type Employ struct {
	Id   string
	Name string
	Age  int
}

//方法 不加*会进行内存拷贝
func (e *Employ) String() string {
	fmt.Printf("address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("id %s name %s age %d", e.Id, e.Name, e.Age)
}

func TestObject(t *testing.T) {
	e := Employ{"0", "Tony", 20}
	e1 := Employ{Id: "2", Name: "Lio"}
	e2 := new(Employ)
	e2.Id = "3"
	e2.Name = "Kiven"
	t.Log(e)
	t.Log(e1)
	t.Log(e2)              //输出&{3 Kiven 0}
	t.Logf("e is %T", &e)  //输出实例类型,加上&为指针类型
	t.Logf("e2 is %T", e2) //输出实例类型:指针类型
	t.Log(e.String())
}
