package data

import (
	"github.com/samurai-of-honor/game-of-life/internal/core"
	"strconv"
	"strings"
)

func ParseInput(str string) (int, core.Universe) {
	str = strings.ReplaceAll(str, "\r", "")
	lines := strings.Split(str, "\n")
	iteration, _ := strconv.Atoi(lines[0])

	universe := strings.Join(lines[2:], "\n")
	return iteration, core.ParseUniverse(universe)
}
