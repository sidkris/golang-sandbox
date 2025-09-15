package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	frontseatPassenger Passenger
}

func main() {
	casey := Passenger{Name: "Casey", TicketNumber: 1, Boarded: false}
	bill := Passenger{Name: "Bill", TicketNumber: 2, Boarded: true}
	marie := Passenger{Name: "Marie", TicketNumber: 3, Boarded: false}
	fmt.Println(casey, bill, marie)
}
