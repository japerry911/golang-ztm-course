package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	Frontseat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println("casey", casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println("bill", bill, "ella", ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println("heidi", heidi)

	heidi.Boarded = true
	bus := Bus{Frontseat: heidi}
	fmt.Println("bus", bus)
}
