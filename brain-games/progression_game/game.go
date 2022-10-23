package proggame

import (
	"fmt"
	braingames "hexletcode/brain-games"
	rand "math/rand"
	"strings"
)

const (
	startMin    = 1
	startMax    = 20
	stepMin     = 2
	stepMax     = 10
	lenMin      = 6
	lenMax      = 12
	filler      = ".."
	min         = 1
	max         = 100
	description = "What number is missing in the progression?"
)

func makeProgresison(start, step, len int) []int {
	result := []int{}
	elem := start
	for i := 0; i <= len; i++ {
		elem += step
		result = append(result, elem)
	}
	return result
}

func game() map[string]string {
	braingames.SetSeed()
	start := rand.Intn(startMax) + startMin
	step := rand.Intn(stepMax) + stepMin
	len := rand.Intn(lenMax) + lenMin

	progression := makeProgresison(start, step, len)
	hiddenNumber := progression[rand.Intn(len)]

	var b strings.Builder
	for _, elem := range progression {
		if elem == hiddenNumber {
			b.WriteString(" " + filler)
		} else {
			b.WriteString(" " + fmt.Sprint(elem))
		}
	}
	question := strings.Trim(b.String(), " ")
	answer := fmt.Sprintf("%d", hiddenNumber)

	gameData := make(map[string]string)
	gameData["question"] = question
	gameData["correctAnswer"] = answer
	return gameData
}

func Run() {
	braingames.Engine(game, description)
}
