package main

import "fmt"
type person struct {
	firstName string
	lastName string
	contact contactInfo
}
type contactInfo struct {
	email string
	zipcode int
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	jim := person {
		firstName: "Jim",
		lastName: "Daniel",
		contact: contactInfo {
			email: "jim@gmail.com",
			zipcode: 20090,
		},
	}

	fmt.Printf("%+v", alex)
	fmt.Printf("%+v", jim)
}