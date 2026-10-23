package main

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
}
