package main

import "fmt"

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
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "Jim@gmail.com",
			zipCode: 16132,
		},
	}
	jim.updateName("Sergio")
	jim.Print()

	carl := person{
		firstName: "Carl",
		lastName:  "Roberts",
		contact: contactInfo{
			email:   "Carl@gmail",
			zipCode: 16133,
		},
	}

	carl.updateName("Sancio")
	carl.Print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) Print() {
	fmt.Printf("%+v", p)
}
