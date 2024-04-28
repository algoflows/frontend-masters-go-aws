package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) changeName(newName string) {
	fmt.Println("Address of copy", &p.Name)
	p.Name = newName
}

func main() {
	// create new person
	myPerson := NewPerson("Sean", 36)
	fmt.Println("Address of allocated memory", &myPerson.Name)
	fmt.Printf("Person before change %+v\n", myPerson)

	a := 7
	b := &a // we get the memory address of a and assign it to b
	*b = 9  // then we point to address of a which is stored in b and point to the value and put 9 in there

	// a becomes 9
	fmt.Println("what is b?", a)

	// change the name
	myPerson.changeName("Melky")

	// log the final person 
	fmt.Printf("Person after change %+v\n", myPerson)

	mySlice := []int{
		1, 2, 3,
	}

	for index, _ := range mySlice {
		mySlice[index]++
	}

	fmt.Println(mySlice)
}
