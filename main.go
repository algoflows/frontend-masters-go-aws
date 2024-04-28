package main

import (
	"fmt"
	"github.com/algoflows/frontend-masters-go-aws/imports"
)

func main() {
	fmt.Println("Hello, world!")

	newTicket := imports.Ticket{
		ID:    123,
		Event: "Money Balling",
	}

	newTicket.PrintEvent()
	fmt.Println(newTicket)
}
