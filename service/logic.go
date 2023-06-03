package service

//package needs to be defined
import (
	"fmt"
)

func WelcomeMessage(tickets uint, remainingTickets uint) {

	fmt.Println("Welcome to our booking portal")
	fmt.Printf("There are %v tickets in total\n", tickets)
	fmt.Printf("%v tickets remaning\n", remainingTickets)
	fmt.Printf("\n")
	fmt.Println("Book your tickets here.")
	fmt.Printf("\n")

}
