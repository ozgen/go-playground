package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	name    string
	surname string
	contact contactInfo
}

func main() {
	//p := person{name: "test", surname: "testSurname"}
	//fmt.Println(p)
	//var alex person
	//alex.name = "alex"
	//alex.surname = "anderson"
	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)
	jim := person{name: "jim", surname: "party",
		contact: contactInfo{
			email:   "jim@email.com",
			zipCode: 123,
		},
	}
	//fmt.Printf("%+v", jim)
	jim.print()
	//jimPointer := &jim
	//jimPointer.updateName("jimmy")
	jim.updateName("jimmy2")
	jim.print()
}

func (p *person) updateName(name string) {
	p.name = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
