package main

import (
	"github.com/samurai-of-honor/game-of-life/internal/core"
	"github.com/samurai-of-honor/game-of-life/internal/data"
)

func main() {
	inputStr := core.ReadFile("./file/startUniverse.txt")
	nTimes, universe := data.ParseInput(inputStr)

	finalUniverse := universe.Evolve(nTimes)
	core.WriteFile(finalUniverse)
}
