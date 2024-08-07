package main

import "fmt"

const LINE = "##############################"

type Person struct {
	name           string
	age            int
	height, weight float32
}

func (person Person) getName() string {
	return person.name
}

func (person Person) getAge() int {
	return person.age
}

func (person Person) getHeight() float32 {
	return person.height
}

func (person Person) getWeight() float32 {
	return person.weight
}

func (person Person) setName1(name string) {
	person.name = name
}

func (person *Person) setName2(name string) {
	person.name = name
}

func GET_INFO(person Person) {
	fmt.Println(LINE)
	fmt.Println("name:", person.name)
	fmt.Println("age:", person.age)
	fmt.Println("height:", person.height)
	fmt.Println("weight:", person.weight)
	fmt.Println()
}

func main() {
	//person1 := Person{name: "Aleksander", age: 21, height: 186.6, weight: 64.5}
	//GET_INFO(person1)
	//person1.setName1("Dima Godyna")
	//GET_INFO(person1)
	//person1.setName2("Dima Godyna")
	//GET_INFO(person1)
	//
	//person2 := &person1
	//person2.setName2("Oleg")
	//GET_INFO(person1)
	//GET_INFO(*person2)

}
