package main

import (
	"fmt"
)

type person struct {
	fname   string
	lname   string
	favfood string
}

func main() {
	p1 := person{"Vincent", "Bailey", "Sandwiches"}
	fmt.Println(p1)

}
