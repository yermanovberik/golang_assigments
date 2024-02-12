package models

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func (p *Person) Print() {
	fmt.Println(p.Name, p.Age, p.City)
}
