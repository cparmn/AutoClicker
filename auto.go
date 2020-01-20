package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	var mleft bool
	for !mleft {
		fmt.Println("Click on Location")
		mleft = robotgo.AddEvent("mleft")
	}
	if mleft {
		x, y := robotgo.GetMousePos()
		robotgo.MoveMouse(x, y)
		for i, r := x, y; i == x; x, y = robotgo.GetMousePos() {
			ClickTheButton(i, r)
		}
	}
}
func delaySecond(n time.Duration) {
	time.Sleep(n * time.Second / 3)
}

func ClickTheButton(x, y int) {
	robotgo.MoveMouseSmooth(x, y, 1.0, 7.0)
	robotgo.Click()
	delaySecond(1)
	robotgo.Click()
}
