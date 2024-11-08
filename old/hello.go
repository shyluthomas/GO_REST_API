package hello

import (
	"fmt"

	"example.com/greetings"
)

func hello() {
	prop := []string{"shylu", "steffy"}
	message, err := greetings.HelloMessage(prop)
	if err == nil {
		fmt.Println("Hello, World!", message)
	} else {
		fmt.Println(err)
	}
}
