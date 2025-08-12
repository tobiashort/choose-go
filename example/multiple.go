package main

import (
	"fmt"

	"github.com/tobiashort/choose-go"
)

func main() {
	options, ok := choose.Multiple("For my smoothie I choose:", []string{
		"Apple",
		"Banana",
		"Lemon",
		"Dragonfruit",
	})
	if ok {
		fmt.Printf("%v\n", options)
	} else {
		fmt.Println("Abort.")
	}
}
