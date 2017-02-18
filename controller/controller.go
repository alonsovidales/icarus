package main

import (
	"fmt"

	"github.com/alonsovidales/icarus/flight"
	"github.com/splace/joysticks"
)

func main() {
	fc := flight.New(17, 27, 22, 18, 23, 2000, 1000, "/dev/pigpio", true)

	events := joysticks.Capture(
		joysticks.Channel{1, joysticks.HID.OnPanX}, // event[4] hat #1 rotates
		joysticks.Channel{1, joysticks.HID.OnPanY}, // event[5] hat #2 rotates
		joysticks.Channel{2, joysticks.HID.OnPanX}, // event[4] hat #1 rotates
		joysticks.Channel{3, joysticks.HID.OnPanY}, // event[5] hat #2 rotates
	)
	fmt.Println(events)
	for {
		select {
		case e := <-events[0]:
			perc := int(e.(joysticks.HatPanXEvent).V * 100)
			if perc > 0 {
				fc.TurnRight(perc)
			} else {
				fc.TurnLeft(-1 * perc)
			}
			fmt.Println("Left X: ", perc)
		case e := <-events[1]:
			perc := int(e.(joysticks.HatPanYEvent).V * 100)
			if perc <= 0 {
				fc.SetTrottle(-1 * perc)
			}
			fmt.Println("Left Y: ", e.(joysticks.HatPanYEvent).V)
		case e := <-events[2]:
			perc := int(e.(joysticks.HatPanXEvent).V * 100)
			if perc > 0 {
				fc.MoveRight(perc)
			} else {
				fc.MoveLeft(-1 * perc)
			}
			fmt.Println("Right X: ", e.(joysticks.HatPanXEvent).V)
		case e := <-events[3]:
			perc := int(e.(joysticks.HatPanYEvent).V * 100)
			if perc > 0 {
				fc.Front(perc)
			} else {
				fc.Back(-1 * perc)
			}
			fmt.Println("Right Y: ", e.(joysticks.HatPanYEvent).V)
		}
	}
}
