package main

import (
	"fmt"

	"github.com/mcclayac/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	// Type Assertion
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a Taper Recorder")
	}
	// END - Type Assertions

	MainHorn()

	ToggleMain()

	AssertionMain()

	ErrorMain()

	StringgerMain()
}
