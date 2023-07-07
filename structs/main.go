package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName:   "Jim",
		lastName:    "Party",
		contactInfo: contactInfo{email: "jim.party@gmail.com", zipCode: 13000},
	}

	//Memory address of jim
	jimPointer := &jim
	fmt.Printf("Memory Address of Jim %p", jimPointer)
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	// fmt.Printf("%+v", p)
}
