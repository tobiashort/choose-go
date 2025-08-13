package main

import (
	"fmt"

	"github.com/tobiashort/choose-go"
)

func main() {
	options, ok := choose.ManyN("For my smoothie I choose:", []string{
		"Apple",
		"Banana",
		"Lemon",
		"Dragonfruit",
	}, 2)
	if ok {
		fmt.Printf("%v\n", options)
	} else {
		fmt.Println("Abort.")
	}
}
