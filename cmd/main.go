package main

import (
	"github.com/samurai-of-honor/game-of-life/internal/core"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func main() {
	width, height, _ := terminal.GetSize(int(os.Stdout.Fd()))

	randomUn := core.GenerateRandomUniverse(width-1, height-2)
	randomUn.EvolveInf(200)
}
