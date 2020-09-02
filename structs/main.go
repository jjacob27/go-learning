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
	p1 := person{
		firstName: "jiju",
		lastName:  "jacob",
		contact: contactInfo{
			email: "j@j.com", zipCode: 12345}}
	p1.updateName("Eric")
	p1.print()
}

func (p *person) updateName(n string) {
	p.firstName = n
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
