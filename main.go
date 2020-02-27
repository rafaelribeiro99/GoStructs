package main

import "fmt"

type contactInfo struct {
	email    string
	zipCoode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)

	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"

	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:    "jim@gmail.com",
			zipCoode: 94000,
		},
	}

	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
