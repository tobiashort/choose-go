package main

import (
	"fmt"

	"github.com/tobiashort/choose-go"
)

func main() {
	idx, option, ok := choose.Single([]string{
		"Option 1",
		"Option 2",
		"Option 3",
	})
	if !ok {
		fmt.Println("Abort.")
	} else {
		fmt.Printf("You have chosen: [%d] %s\n", idx, option)
	}
}
