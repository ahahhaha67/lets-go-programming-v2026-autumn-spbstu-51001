package main

import "fmt"

func main() {

	var firstArg int64
	_, err := fmt.Scan(&firstArg)
	if err != nil {
		fmt.Println("error reading first argument")
		return
	}
	var secondArg int64
	_, err = fmt.Scan(&secondArg)
	if err != nil {
		fmt.Println("error reading second argument")
		return
	}
	var operator string
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("error reading operator")
		return
	}

	switch operator {
	case "+":
		fmt.Print(firstArg + secondArg)
	case "-":
		fmt.Print(firstArg - secondArg)
	case "*":
		fmt.Print(firstArg * secondArg)
	case "/":
		if secondArg == 0 {
			fmt.Println("can't divide by zero")
		} else {
			fmt.Println(firstArg / secondArg)
		}
	default:
		fmt.Println("unknown operator")
	}
}
