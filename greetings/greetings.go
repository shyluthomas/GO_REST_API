package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

type Props struct {
	Name string
}

// Hello returns a greeting for the named person.
func Hello(prop Props) (string, error) {
	log.SetPrefix("greeting:")
	log.SetFlags(0)
	if prop.Name == "" {
		log.Fatal("name is empty")
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), prop.Name)
	return message, nil
}

func randomFormat() string {
	format := []string{
		"hi %v, welocme",
		"great to see you %v",
		"Hal hello sir %v",
	}
	return format[rand.Intn(len(format))]
}

func HelloMessage(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		prop := Props{Name: name}
		message, err := Hello(prop)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
