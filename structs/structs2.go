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
	casey := Passenger{Name: "Casey", TicketNumber: 1, Boarded: true}
	bill := Passenger{Name: "Bill", TicketNumber: 2, Boarded: true}
	marie := Passenger{Name: "Marie", TicketNumber: 3, Boarded: false}
	fmt.Println(casey, bill, marie)

	if casey.Boarded {
		fmt.Println("Casey has boarded the bus.")
	}

	if bill.Boarded {
		fmt.Println("Bill has boarded the bus.")
	}

	if marie.Boarded {
		fmt.Println("Marie has boarded the bus.")
	}

	bus := Bus{frontseatPassenger: casey}
	fmt.Println(bus.frontseatPassenger.Name, "is the front seat passenger.")
}
