package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

func main() {
	p1 := person{"Vincent", "Bailey"}
	fmt.Println(p1.fname)
}
