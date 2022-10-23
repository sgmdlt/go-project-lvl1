package main

import (
	"fmt"
	braingames "hexletcode/brain-games"
	calcgame "hexletcode/brain-games/calc_game"
	evengame "hexletcode/brain-games/even_game"
	gcdgame "hexletcode/brain-games/gcd_game"
)

func main() {
	fmt.Println("Please enter the game number and press Enter.")
	fmt.Println("0 - Exit")
	fmt.Println("1 - Greet")
	fmt.Println("2 - Even")
	fmt.Println("3 - Calc")
	fmt.Println("4 - GCD")
	var choice string
	fmt.Scanln(&choice)
	fmt.Printf("Your choice: %s\n", choice)
	switch choice {
	case "0":
		return
	case "1":
		braingames.WelcomeUser()
	case "2":
		evengame.Run()
	case "3":
		calcgame.Run()
	case "4":
		gcdgame.Run()
	}
}
