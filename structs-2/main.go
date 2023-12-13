package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int64
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// fmt.Println("Hello, World!")
	jim := person{firstName: "Jim", lastName: "Nelson", contact: contactInfo{
		email:   "jim@nelson.org",
		zipCode: 01203,
	}}

	jimPointer := &jim

	jim.print()
	jimPointer.updateJim("Jimmie")
	fmt.Println("-------------------------------")
	jim.print()
}

func (p *person) updateJim(firstName string) {
	p.firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
