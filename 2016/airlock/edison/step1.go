package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/intel-iot/edison"
)

func main() {
	board := edison.NewAdaptor()
	blue := gpio.NewGroveLedDriver(board, "3")

	work := func() {
		gobot.Every(1*time.Second, func() {
			blue.Toggle()
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{blue},
		work,
	)

	robot.Start()
}
