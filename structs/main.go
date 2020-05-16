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

	var name2 person
	name2.firstName = "Vinay"
	name2.lastName = "Kotegowder"
	name2.contact.email = "vinay.kotegowder@gmail.com"
	name2.contact.zipCode = 577004
	fmt.Printf("%+v", name2)
}
