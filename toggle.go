package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func ToggleMain() {
	fmt.Println("------------------------------------")
	fmt.Println(" ToggleMain()")
	fmt.Println("------------------------------------")
	s := Switch("off")
	// var t Toggleable = s
	var t Toggleable = &s

	t.toggle()
	t.toggle()
	t.toggle()
	t.toggle()
	t.toggle()
	t.toggle()
}
