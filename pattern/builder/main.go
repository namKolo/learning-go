package main

import (
	"fmt"
	"learning-go/pattern/builder/car"
	"learning-go/pattern/builder/messenger"
)

func main() {
	builder := car.NewBuilder()
	car := builder.TopSpeed(50).Paint(car.BLUE).Build()

	fmt.Println(car.Drive())
	fmt.Println(car.Stop())

	sender := &messenger.Sender{}

	jsonMsg, err := sender.BuildMessage(&messenger.JSONMessageBuilder{})

	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonMsg.Body))

	xmlMsg, err := sender.BuildMessage(&messenger.XMLMessageBuilder{})

	if err != nil {
		panic(err)
	}

	fmt.Println(string(xmlMsg.Body))
}
