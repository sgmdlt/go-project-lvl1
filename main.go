package main

import (
	"fmt"
	braingames "hexletcode/brain-games"
	calcgame "hexletcode/brain-games/calc_game"
	evengame "hexletcode/brain-games/even_game"
)

func main() {
	fmt.Println("Please enter the game number.")
	fmt.Println("0 - Exit")
	fmt.Println("1 - Greet")
	fmt.Println("2 - EvenGame")
	fmt.Println("3 - CalcGame")
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
	}
}
