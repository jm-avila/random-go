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

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// In the context a variable
// & = value to pointer eg. thomasPointer := &thomas
// * = pointer to value eg. thomasValue := *thomasPointer

// In the context a type
// * = the value is a pointer of a type

func main() {
	// Example 1
	// thomas := person{firstName: "Thomas", lastName: "Anderson"}
	// fmt.Println("thomas")
	// fmt.Println(thomas)

	// Example 2
	// var alex person
	// fmt.Println("alex")
	// fmt.Printf("%+v\n", alex)
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)

	// Example 3
	jimmy := person{
		firstName: "Mike",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "mike@mail.com",
			zipCode: 94000,
		}}
	fmt.Println("jimmy")
	jimmy.print()
	jimmy.updateName("Jimmy")
	jimmy.print()
}
