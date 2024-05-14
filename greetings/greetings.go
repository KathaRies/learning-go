package greetings

import (
	"errors"
	"fmt"

	"math/rand"

	"rsc.io/quote"
)

func Hello(name string, time int) (string, error) {
	if name == "" {
		{
			return "", errors.New("empty name")
		}
	}
	formats := []string{
		fmt.Sprintf("%v, %v", quote.Opt(), name),
		fmt.Sprintf("%v, %v", formalGreeting(time), name),
		fmt.Sprintf("Hi, %v. Welcome to this little world", name),
	}

	return formats[rand.Intn(len(formats))], nil
}

func formalGreeting(time int) string {
	if time < 12 {
		return "Good morning"
	}
	if time > 12 && time < 18 {
		return "Good afternoon"
	}
	if time > 18 {
		return "Good evening"
	}
	return "Good day"
}

func Hellos(names []string, time int) (map[string]string, error) {
	msgs := make(map[string]string)
	for _, name := range names {
		msg, err := Hello(name, time)
		if err != nil {
			return nil, err
		}
		msgs[name] = msg
	}
	return msgs, nil
}
