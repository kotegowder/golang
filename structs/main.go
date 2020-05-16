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
	// Has to be put in order of members
	//name1 := person{"Vinay", "Kotegowder"}
	// Declared against property name
	//name1 := person{firstName: "Vinay", lastName: "Kotegowder"}
	var name1 person
	// Automatically asign empty string
	// fmt.Println(name1)
	// Prints along with the struct members
	// fmt.Printf("%+v", name1)
	name1.firstName = "Vinay"
	name1.lastName = "Kotegowder"
	fmt.Printf("%+v", name1)

	name2 := person{
		firstName: "Vinay",
		lastName:  "Kotegowder",
		contactInfo: contactInfo{
			email:   "vinay.kotegowder@gmail.com",
			zipCode: 577004,
		},
	}

	name2.details()
	//name2Pointer := &name2
	//name2Pointer.updateName("Vidhya")
	name2.updateName("Vidhya")
	name2.details()
}

func (p person) details() {
	fmt.Printf("%+v", p)
}

// *person -> This is a type declaration - it means we're working with a pointer to a person
// *pToPerson -> This is an operator - it means we want to manipualate the value the pointer is referencing
func (pToPerson *person) updateName(firstName string) {
	(*pToPerson).firstName = firstName
}
