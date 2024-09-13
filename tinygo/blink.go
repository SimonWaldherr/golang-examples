package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.Pin(25) // Use Pin 19 on Pico W
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.High()
		time.Sleep(time.Millisecond * 500)
		led.Low()
		time.Sleep(time.Millisecond * 500)
	}
}
