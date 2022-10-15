package main

import (
	"ip-0x-lab-1-samurai-of-honor/internal/data"
)

func main() {
	inputStr := data.ReadFile("./cmd/startUniverse.txt")
	nTimes, universe := data.ParseInput(inputStr)
	finalUniverse := universe.Evolve(nTimes)
	data.WriteFile(finalUniverse)
}
