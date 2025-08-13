package main

import (
	"fmt"

	"github.com/tobiashort/choose-go"
)

func main() {
	idx, _, ok := choose.One("Choose your poison:", []string{
		"Venomous Snake Milkshake",
		"Radioactive Sushi Platter",
		"Carolina Reaper Lava Stew",
		"Brain Smoothie à la Zombie",
	})
	if !ok {
		fmt.Println("You survived!")
	} else {
		switch idx {
		case 0:
			fmt.Println("You drank the snake milkshake. It bit back.")
		case 1:
			fmt.Println("You ate the glowing sushi. Now you glow too.")
		case 2:
			fmt.Println("You slurped lava stew. Your tongue quit its job.")
		case 3:
			fmt.Println("You sipped the brain smoothie. The zombie says thanks.")
		}
	}
}
