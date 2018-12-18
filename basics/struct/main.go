package main

import "fmt"

type contactInfo struct {
	emailAddress string
	zipCode      int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) printName() {
	fmt.Println(p.firstName, p.lastName)
}

func (p *person) updateFirstName(fName string) {
	p.firstName = fName
}

func main() {
	anthony := person{
		firstName: "Anthony",
		lastName:  "Freda",
		contactInfo: contactInfo{
			zipCode:      12345,
			emailAddress: "anthonyfreda323@gmail.com",
		},
	}
	anthony.updateFirstName("Ant")
	anthony.print()
}
