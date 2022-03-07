package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// Go executes init functions automatically at program startup,
// after global variables have been initialized.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat starts with a lowercase letter,
// making it accessible only to code in its own package
func randomFormat() string {
	formats := []string{
		"Hello, %v. Welcome!",
		"Great to see you, %v!",
		"Well met, %v.",
	}

	return formats[rand.Intn(len(formats))]
}
