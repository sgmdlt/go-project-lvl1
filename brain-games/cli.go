package braingames

import (
	"fmt"
)

func WelcomeUser() string {
	fmt.Println("Welcome to the Brain Games!")
	fmt.Println("May I have your name? ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n", name)
	return name
}
