package main

import (
	"fmt"
	"learn/utils/models"
	"reflect"
)

type Printer interface {
	Print()
}

func DisplayInfo(printer Printer) {
	printer.Print()
	fmt.Println()
}

func Map() {
	personMap := map[string]string{}
	personMap["name"] = "Berik"
	personMap["age"] = "20"
	personMap["city"] = "Almaty"
}

func checkEvenOrOdd(number int) {
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func main() {
	var intValue int = 5
	fmt.Println(intValue)
	fmt.Println(reflect.TypeOf(intValue))

	var myName string = "Berik"
	fmt.Println(myName)

	number1, number2 := 3, 4
	checkEvenOrOdd(number1)
	checkEvenOrOdd(number2)

	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:3]
	fmt.Println(slice)

	Map()

	person1 := models.Person{
		Name: "Berik",
		Age:  20,
		City: "Turkestan",
	}

	person2 := models.Person{
		Name: "Didar",
		Age:  21,
		City: "Kyzylorda",
	}

	person1.Print()
	person2.Print()

}
