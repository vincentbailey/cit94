package main

import "fmt"

// 3. type person struct
type person struct {
	fName string
	lName string
	// 5. add favFood to struct
	favFood string
}

func main() {

	// 1. slice of int
	xi := []int{2, 4, 6, 8}
	//range over printing just index
	for i := range xi {
		fmt.Println(i)
	}
	//range over printing index and value
	for i, v := range xi {
		fmt.Println(i, v)
	}

	// 2. intitialize map and print
	m := map[string]int{"this is a map key": 94}
	fmt.Println(m)
	//range over map, print key
	for k := range m {
		fmt.Println(k)
	}
	//range over map, print key and value
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 4. create value type person assign identifier p1, print fname
	// 5. add values to field favFood, print out values, also print using for range loop
	p1 := person{"first name", "last name", "sandwiches, soups, salads"}
	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)
	for i, v := range favFood {
		fmt.Println(i, v)
	}

}
