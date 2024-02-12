package models

import (
	"fmt"
)

type Employee struct {
	Person   Person
	JobTitle string
}

func (e *Employee) Print() {
	fmt.Println(e.Person.Name, e.Person.Age, e.Person.City, e.JobTitle)
}
