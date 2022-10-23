package evengame

import (
	"fmt"
	braingames "hexletcode/brain-games"
	rand "math/rand"
)

const (
	min         = 1
	max         = 100
	description = "Answer 'yes' if the number is even, otherwise answer 'no'."
)

func isEven(num int) bool {
	return num%2 == 0
}

func game() map[string]string {
	braingames.SetSeed()
	num := rand.Intn(max) + min

	question := fmt.Sprintf("%d", num)
	var answer string
	gameData := make(map[string]string)

	if isEven(num) {
		answer = "yes"
	} else {
		answer = "no"
	}

	gameData["question"] = question
	gameData["correctAnswer"] = answer
	return gameData
}

func Run() {
	braingames.Engine(game, description)
}
