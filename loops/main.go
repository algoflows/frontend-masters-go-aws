package main

import (
	"fmt"
	"slices"
)

const (
	PI = 3.14
)

func main() {
	animals := []string{
		"dog",
		"cat",
	}

	animals = append(animals, "moose")
	animals = slices.Delete(animals, 0, 1)

	for i := 0; i < len(animals); i++ {
		fmt.Printf("This my animal %s\n", animals[i])
	}

	for _, animal := range animals {
		fmt.Println(animal)
	}

	for index, value := range animals {
		fmt.Printf("This is my index %d and this is my animal %s\n", index, value)
	}

	for value := range 11 {
		fmt.Println(value)
	}

	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}
}
