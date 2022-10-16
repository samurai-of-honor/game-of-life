package main

import (
	"github.com/samurai-of-honor/game-of-life/internal/data"
)

func main() {
	inputStr := data.ReadFile("./cmd/startUniverse.txt")
	nTimes, universe := data.ParseInput(inputStr)
	finalUniverse := universe.Evolve(nTimes)
	data.WriteFile(finalUniverse)
}
