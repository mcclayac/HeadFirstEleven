package main

import "fmt"

type NoiseMaker interface {
	MakeSound()
}

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!,", w)
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!,", h)
}

// --------------------------
// Take Interface as parameter
func play(n NoiseMaker) {
	n.MakeSound()
}

// --------------------------
func MainHorn() {
	fmt.Println("------------------------------------")
	fmt.Println(" MainHorn()")
	fmt.Println("------------------------------------")
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()

	toy = Horn("Toyco Blaster")
	toy.MakeSound()

	play(Whistle("Toyco Canary 2"))
	play(Horn("Toyco Blaster 2"))

}
