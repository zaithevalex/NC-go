package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	var Person = person{}
	Person.name = name
	Person.age = 20
	return &Person
}

func main() {
	fmt.Println(person{"Bob", 20})

	Person1 := person{"Semen", 21}
	fmt.Println("Person1:", Person1)
	Person1.age = 22
	fmt.Println("Person1:", Person1)

	var Person2 = person{"Arseny", 21}
	fmt.Println("Person2:", Person2)

	var Person3 = &Person2
	fmt.Println("Person3:", Person3)
	Person3.age = 30
	fmt.Println("Person3:", Person3, *Person3)

	Person4 := newPerson("Alex")
	fmt.Println("Person4:", Person4, *Person4)

	// change field in methods of different functions versions:
	Person5 := person{"SWAG", 333}
	fmt.Println("1st step, Person5:", Person5, &Person5)
	changePerson1stVersion(Person5)
	fmt.Println("2nd step, Person5:", Person5, &Person5)
	changePerson2ndVersion(&Person5)
	fmt.Println("3rd step, Person5:", Person5, &Person5)
}

func changePerson1stVersion(Person person) {
	Person.name = "11111"
	Person.age = 111
}

func changePerson2ndVersion(Person *person) {
	Person.name = "11111"
	Person.age = 111
}
