package braingames

import "fmt"

const tries = 3

func Engine(game func() map[string]string, description string) {
	var userAnswer string
	username := WelcomeUser()
	fmt.Println(description)

	for i := 0; i < tries; i++ {
		gameData := game()
		question := gameData["question"]
		correctAnswer := gameData["correctAnswer"]
		fmt.Printf("Question: %s\n", question)

		fmt.Println("Your answer: ")
		fmt.Scanln(&userAnswer)

		if userAnswer != correctAnswer {
			fmt.Printf("%s is wrong answer ;(. Correct answer was %s\n", userAnswer, correctAnswer)
			fmt.Printf("Let's try again, %s!", username)
			return
		}

		fmt.Printf("Correct!\n")
	}
	fmt.Printf("Congratulations, %s", username)
}
