package main

import "fmt"

func main() {

	var firtArg int64
	var secondArg int64
	var operator string

	fmt.Scan(&firtArg, &secondArg, &operator)

	switch operator {
	case "+":
		fmt.Print(firtArg + secondArg)
	case "-":
		fmt.Print(firtArg - secondArg)
	case "*":
		fmt.Print(firtArg * secondArg)
	case "/":
		if secondArg == 0 {
			fmt.Println("can't divide by zero")
		} else {
			fmt.Println(firtArg / secondArg)
		}
	default:
		fmt.Println("unknown operator")
	}
}
