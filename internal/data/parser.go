package data

import (
	"ip-0x-lab-1-samurai-of-honor/internal/core"
	"strconv"
	"strings"
)

func ParseInput(str string) (int, core.Universe) {
	lines := strings.Split(str, "\n")
	iteration, _ := strconv.Atoi(lines[0])

	universe := strings.Join(lines[2:], "\n")
	return iteration, core.ParseUniverse(universe)
}
