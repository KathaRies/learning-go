package main

import (
	"fmt"
	"log"

	"go/greetings"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("greeting s")
	msg, err := greetings.Hello("Sven", 23)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	msgs, _ := greetings.Hellos([]string{"Ana", "Kate", "Tom"}, 8)
	fmt.Println(msgs)
}
