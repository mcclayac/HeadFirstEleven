package main

import (
	"fmt"
	"log"
)

type ComedyError string

func (c ComedyError) Error() string {
	return string(c + ", Theze nutz")
}

// ---------------------------------------
type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func ErrorMain() {
	fmt.Println("------------------------------------")
	fmt.Println(" ErrorMain()")
	fmt.Println("------------------------------------")
	var err error
	err = ComedyError("What's a programmer's favorite beer?  Logger!")
	fmt.Println(err)

	// -----------------------------------------
	err = OverheatError(103.3)
	fmt.Println(err)

	err = checkTemperature(121.379, 200.0)
	if err != nil {
		log.Fatal(err)
	}
}
