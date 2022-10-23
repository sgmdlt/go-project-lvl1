package gcdgame

import (
	"fmt"
	braingames "hexletcode/brain-games"
	rand "math/rand"
)

const (
	min         = 1
	max         = 100
	description = "Find the greatest common divisor of given numbers."
)

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func game() map[string]string {
	braingames.SetSeed()
	first := rand.Intn(max) + min
	second := rand.Intn(max) + min

	question := fmt.Sprintf("%d %d", first, second)
	answer := fmt.Sprintf("%d", gcd(first, second))

	gameData := make(map[string]string)
	gameData["question"] = question
	gameData["correctAnswer"] = answer
	return gameData
}

func Run() {
	braingames.Engine(game, description)
}
