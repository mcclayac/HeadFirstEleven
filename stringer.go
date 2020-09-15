package main

import "fmt"

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func StringgerMain() {
	fmt.Println("------------------------------------")
	fmt.Println(" StringgerMain()")
	fmt.Println("------------------------------------")
	coffeepot := CoffeePot("LuxBrew")

	fmt.Print(coffeepot, "\n")
	fmt.Println(coffeepot)
	fmt.Printf("%s\n\n", coffeepot)

}
