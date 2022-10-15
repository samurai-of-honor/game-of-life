package core

import "strings"

type Universe [][]bool

func (u Universe) getCell(x, y int) bool {
	if x < 0 {
		x += len(u[0])
	} else if x >= len(u[0]) {
		x -= len(u[0])
	}

	if y < 0 {
		y += len(u)
	} else if y >= len(u) {
		y -= len(u)
	}

	return u[y][x]
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (u Universe) searchNeighbors(x, y int) int {
	var count int
	count += boolToInt(u.getCell(x-1, y-1))
	count += boolToInt(u.getCell(x, y-1))
	count += boolToInt(u.getCell(x+1, y-1))
	count += boolToInt(u.getCell(x-1, y))
	count += boolToInt(u.getCell(x+1, y))
	count += boolToInt(u.getCell(x-1, y+1))
	count += boolToInt(u.getCell(x, y+1))
	count += boolToInt(u.getCell(x+1, y+1))
	return count
}

func (u Universe) NextStep() Universe {
	newUn := make(Universe, len(u))
	for i := range u {
		newUn[i] = make([]bool, len(u[i]))
	}

	for y, line := range u {
		for x, cell := range line {
			switch u.searchNeighbors(x, y) {
			case 3:
				newUn[y][x] = true
			case 2:
				newUn[y][x] = cell
			}
		}
	}

	return newUn
}

func (u Universe) Evolve(nTimes int) Universe {
	for i := 0; i < nTimes; i++ {
		u = u.NextStep()
	}
	return u
}

func (u Universe) ToString() string {
	var resultStr string

	for _, line := range u {
		for _, cell := range line {
			if cell == true {
				resultStr += "x"
			} else {
				resultStr += "."
			}
		}
		resultStr += "\n"
	}

	return resultStr
}

func ParseUniverse(str string) Universe {
	lines := strings.Split(str, "\n")
	universe := make(Universe, len(lines))

	for y, line := range lines {
		universe[y] = make([]bool, len(line))

		for x, cell := range line {
			universe[y][x] = cell == 'x'
		}
	}

	return universe
}
