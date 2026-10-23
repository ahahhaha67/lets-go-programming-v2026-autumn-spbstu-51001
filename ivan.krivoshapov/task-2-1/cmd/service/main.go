package main

import "fmt"

const (
	lowestTemp  = 15
	highestTemp = 30
)

func pickTemperature(sign string, value int, from, to *int) int {
	switch sign {
	case ">=":
		if value > *from {
			*from = value
		}
	case "<=":
		if value < *to {
			*to = value
		}
	}

	if *from > *to {
		return -1
	}

	return *from
}

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Invalid departments count")

		return
	}

	for d := 0; d < departments; d++ {
		var staff int
		if _, err := fmt.Scan(&staff); err != nil {
			fmt.Println("Invalid employees count")

			return
		}

		from, to := lowestTemp, highestTemp

		for e := 0; e < staff; e++ {
			var sign string
			var value int

			if _, err := fmt.Scan(&sign, &value); err != nil {
				fmt.Println("Invalid input")

				return
			}

			if sign != ">=" && sign != "<=" {
				fmt.Println("Invalid operator")

				return
			}

			fmt.Println(pickTemperature(sign, value, &from, &to))
		}
	}
}
