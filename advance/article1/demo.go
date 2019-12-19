package article1

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object")
}

func main() {
	flag.Parse()
	fmt.Printf("hello, %s1\n", name)
}
