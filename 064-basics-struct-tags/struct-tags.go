package main

import (
	"fmt"
	"reflect"
)

type Messanger interface {
	GetMessage() string
}

type ErrorMessage struct {
	Message string `prefix:">>>>>>>>" postfix:"<<<<<<<<<<"`
}

func (message ErrorMessage) GetMessage() string {
	return "ERROR: Booo! " + message.Message
}

type LogMessage struct {
	Message string `prefix:"\x1b[32m" postfix:"\x1b[0m"`
}

func (message LogMessage) GetMessage() string {
	return message.Message
}

func show(input Messanger) {
	st := reflect.TypeOf(input)
	field := st.Field(0)

	fmt.Printf(
		"%s%s%s\n",
		field.Tag.Get("prefix"),
		input.GetMessage(),
		field.Tag.Get("postfix"),
	)
}

func main() {
	show(LogMessage{"Hoł hoł"})
	show(ErrorMessage{Message: "Noł noł noł"})
}
