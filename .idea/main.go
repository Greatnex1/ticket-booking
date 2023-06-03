package main

import (
	"fmt"
	"ticket-booking-app/service"
)

const ticket = 50

var remainingTicket uint = 50

func main() {

	service.WelcomeMessage(ticket, remainingTicket)

	fmt.Println()
}
