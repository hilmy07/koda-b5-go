package internals

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name    string
	Address string
	Phone   string
}

func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

func (p Person) Print() string {
	return fmt.Sprintf(
		"Name: %s, Address: %s, Phone: %s",
		p.Name,
		p.Address,
		p.Phone,
	)
}

func (p Person) Greet() {
	fmt.Println("Hello,", p.Name)
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func Method() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter your city: ")
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	fmt.Print("Enter your phone: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	person := NewPerson(name, city, phone)

	fmt.Println(person.Print())
	person.Greet()

	person.SetName("Muhammad Hilmy")
	person.Greet()
}
