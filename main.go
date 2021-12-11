package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	//"golang.design/x/hotkey"
	"github.com/go-vgo/robotgo"
)

func main() {
	//delay := flag.Int("delay", 10, "delay between clicks in miliseconds")
	//flag.Parse()
	var button string = os.Args[1]
	delay, err := strconv.Atoi(os.Args[2])
	if err == nil {
		fmt.Printf("%v", err)
	}
	var forever bool = true
	//var i uint64 = 0
	//i < 99999
	for forever {
		x, y := robotgo.GetMousePos()
		//if the user isn't trying to quit
		if x > 21 && y > 21 {
			time.Sleep(time.Duration(delay) * time.Millisecond)
			robotgo.Click(button, true)
		} else {
			return
		}
		//i++
	}
}
