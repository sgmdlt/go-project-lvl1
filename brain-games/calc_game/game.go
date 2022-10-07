package calc_game

import (
	"fmt"
	braingames "hexletcode/brain-games"
	rand "math/rand"
)

const (
	min         = 1
	max         = 20
	description = "What is the result of the question?"
)

var operations = map[string]func(int, int) int{
	"+": func(a int, b int) int { return a + b },
	"-": func(a int, b int) int { return a - b },
	"*": func(a int, b int) int { return a * b },
}

func calculate(op string, first int, second int) int {
	return operations[op](first, second)
}

func game() map[string]string {
	braingames.SetSeed()
	firstnum := rand.Intn(max) + min
	secondnum := rand.Intn(max) + min
	keys := []string{}
	for key := range operations {
		keys = append(keys, key)
	}
	operation := keys[rand.Intn(len(keys))]

	question := fmt.Sprintf("%d %s %d", firstnum, operation, secondnum)
	answer := fmt.Sprintf("%d", calculate(operation, firstnum, secondnum))

	gameData := make(map[string]string)
	gameData["question"] = question
	gameData["correctAnswer"] = answer
	return gameData
}

func Run() {
	braingames.Engine(game, description)
}
