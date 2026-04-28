package minitask7

import "fmt"

type Person struct {
	Name        string
	Address     string
	PhoneNumber string
}

func NewPerson(name string, address string, phone string) *Person {
	return &Person{
		Name:        name,
		Address:     address,
		PhoneNumber: phone,
	}
}

func (u *Person) Print() {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Address: %s\n", u.Address)
	fmt.Printf("Phone Number: %s\n", u.PhoneNumber)
}

func (u *Person) Greet() {
	fmt.Printf("Hello, %s!\n", u.Name)
}

func (u *Person) SetName(name string) {
	u.Name = name
}
