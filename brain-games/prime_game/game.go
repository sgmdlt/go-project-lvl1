package primegame

import (
	"fmt"
	braingames "hexletcode/brain-games"
	rand "math/rand"
)

const (
	min         = 1
	max         = 100
	description = "Answer 'yes' if given number is prime. Otherwise answer 'no'."
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for div := 2; div <= (num/2)+1; div++ {
		if num%div == 0 {
			return false
		}
	}
	return true
}

func game() map[string]string {
	braingames.SetSeed()
	num := rand.Intn(max) + min

	question := fmt.Sprintf("%d", num)
	var answer string
	if isPrime(num) {
		answer = "yes"
	} else {
		answer = "no"
	}

	gameData := make(map[string]string)
	gameData["question"] = question
	gameData["correctAnswer"] = answer
	return gameData
}

func Run() {
	braingames.Engine(game, description)
}
