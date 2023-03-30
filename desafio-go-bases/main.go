package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	tickets.ReadDataFile("tickets.csv")
	total, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}
	println(total)
	total, _ = tickets.GetMornings("mañana")
	println(total)
	total, _ = tickets.GetMornings("madrugada")
	println(total)
	total, _ = tickets.GetMornings("tarde")
	println(total)
	total, _ = tickets.GetMornings("noche")
	println(total)
	totales, error := tickets.AverageDestination("Brazil")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(totales)
	//fmt.Println(tickets.Tickets)
}
