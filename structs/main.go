package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex", lastName: "Anderson", contact: contactInfo{email: "info@gmail.com", zipCode: 123}}

	alex.updateName("Jhon")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(n string) {
	p.firstName = n
}
