package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id      int
	Name    string
	Email   string
	Country string
	Date    int
	Price   int
}

var Tickets = []Ticket{}

func ReadDataFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		var ticket Ticket
		if value, err := strconv.Atoi(line[0]); err == nil {
			ticket.Id = value
		}
		ticket.Name = line[1]
		ticket.Email = line[2]
		ticket.Country = line[3]
		if value, err := strconv.Atoi(strings.Split(line[4], ":")[0]); err == nil {
			ticket.Date = value
		}
		if value, err := strconv.Atoi(line[5]); err == nil {
			ticket.Price = value
		}
		Tickets = append(Tickets, ticket)
	}
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	if len(Tickets) <= 0 {
		return 0, errors.New("Not found data")
	}
	total := 0
	for _, v := range Tickets {
		if v.Country == destination {
			total++
		}
	}
	return total, nil
}

// ejemplo 2
func GetCountByPeriod(time string) (int, error) {
	if len(Tickets) <= 0 {
		return 0, errors.New("Not found data")
	}
	
	switch time {
	case "madrugada":
		return EvaluateTime(0, 6), nil
	case "mañana":
		return EvaluateTime(7, 12), nil
	case "tarde":
		return EvaluateTime(13, 19), nil
	case "noche":
		return EvaluateTime(20, 23), nil
	default:
		return 0, errors.New("Not found time")
	}
}

func EvaluateTime(from, to int) (total int){
	for _, v := range Tickets {
		if int(v.Date) >= from && v.Date <= to {
			total++
		}
	}
	return total

}

// ejemplo 3
func AverageDestination(destination string) (float64, error) {
	totalCountryTickets, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	if len(Tickets) <= 0 {
		return 0, errors.New("cannot divide by 0")
	}

	return float64(totalCountryTickets) * 100 / float64(len(Tickets)), nil
}
