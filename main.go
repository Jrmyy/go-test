package main

import (
	"github.com/Jrmyy/go-test/services"
	"github.com/Jrmyy/go-test/yolo"
	"strconv"
)

func main() {
	services.DisplayText("Hello World")
	myInt := 1
	myInt = yolo.DoThis(myInt)
	services.DisplayText(strconv.Itoa(myInt))
}
