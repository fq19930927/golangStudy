package extend

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("speak", host)
}

/*type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Println("wang!")
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println("speak", host)
}*/
//内嵌无法当做继承使用
type Dog struct {
	Pet
}

func TestDog(t *testing.T)  {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("lalala")
}
