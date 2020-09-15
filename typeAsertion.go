package main

import "fmt"

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func AssertionMain() {
	fmt.Println("------------------------------------")
	fmt.Println(" AssertionMain()")
	fmt.Println("------------------------------------")
	var noiseMaker NoiseMaker = Robot("Botco Ambler")
	noiseMaker.MakeSound()

	// Type assertion
	var robot Robot = noiseMaker.(Robot)
	robot.Walk()
}
