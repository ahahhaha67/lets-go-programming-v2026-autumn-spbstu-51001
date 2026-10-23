package main

import "fmt"

const (
	lowestTemp  = 15
	highestTemp = 30
)

func pickTemperature(sign string, value int, from, to int) int {
	switch sign {
	case ">=":
		if value > from {
			from = value
		}
	case "<":
		if value < to {
			to = value
		}
	}
	if from > to {
		return -1
	}
	return from
}

func main() {
	var department int
	if _, err := fmt.Scan(&department); err != nil {
		fmt.Print("Invalid department count")

		return
	}

	for d := 0; d < department; d++ {
		var staff int
		if _, err := fmt.Scan(&staff); err != nil {
			fmt.Print("Invalid staff count")

			return
		}

		from, to := lowestTemp, highestTemp

		for e := 0; e < staff; e++ {
			var sign string
			var value int

			if _, err := fmt.Scan(&sign, &value); err != nil {
				fmt.Print("Invalid input")
				return
			}

			if sign != ">=" && sign != "<" {
				fmt.Print("Invalid operator")
				return
			}

			fmt.Print(pickTemperature(sign, value, from, to))
		}

	}

}
