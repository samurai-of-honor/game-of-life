package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleUniverse = Universe{
	{false, false, false, false},
	{false, true, false, false},
	{false, true, false, false},
	{false, true, false, false},
	{false, false, false, false},
}

var u1 = Universe{
	{true, false, false, false, true, true},
	{false, true, true, false, false, false},
	{false, false, true, false, true, false},
}

func TestGetCellExample(t *testing.T) {
	assert.False(t, exampleUniverse.getCell(0, 0))
	assert.False(t, exampleUniverse.getCell(-1, 0))
	assert.False(t, exampleUniverse.getCell(-1, -1))
	assert.False(t, exampleUniverse.getCell(0, -1))

	assert.True(t, exampleUniverse.getCell(1, 2))
	assert.True(t, exampleUniverse.getCell(-3, 3))
	assert.True(t, exampleUniverse.getCell(-3, -3))
}

func TestGetCell1(t *testing.T) {
	assert.False(t, u1.getCell(1, 0))
	assert.False(t, u1.getCell(-3, 2))
	assert.False(t, u1.getCell(-1, -1))
	assert.False(t, u1.getCell(1, 2))

	assert.True(t, u1.getCell(0, 0))
	assert.True(t, u1.getCell(len(u1[0])-2, -1))
	assert.True(t, u1.getCell(2, 2))
}

func TestSearchNeighborsExample(t *testing.T) {
	assert.Equal(t, 3, exampleUniverse.searchNeighbors(2, 2))
	assert.Equal(t, 1, exampleUniverse.searchNeighbors(0, 0))
	assert.Equal(t, 0, exampleUniverse.searchNeighbors(3, 3))
	assert.Equal(t, 0, exampleUniverse.searchNeighbors(-1, -1))
	assert.Equal(t, 2, exampleUniverse.searchNeighbors(0, 3))
	assert.Equal(t, 1, exampleUniverse.searchNeighbors(1, -2))
}

func TestSearchNeighbors1(t *testing.T) {
	assert.Equal(t, 2, u1.searchNeighbors(0, 0))
	assert.Equal(t, 3, u1.searchNeighbors(1, 1))
	assert.Equal(t, 4, u1.searchNeighbors(3, -1))
	assert.Equal(t, 4, u1.searchNeighbors(-1, 2))
	assert.Equal(t, 4, u1.searchNeighbors(len(u1[2])-1, -1))
}

func TestNextStepExample(t *testing.T) {
	wantUniverse := `....
....
xxx.
....
....
`

	outUniverse := exampleUniverse.NextStep()

	assert.Equal(t, wantUniverse, outUniverse.ToString())
}

func TestNextStep1(t *testing.T) {
	inputUniverse := `.......
..x.x..
.x.x.x.
..x.x..
.......`

	wantUniverse := `.......
..xxx..
.x...x.
..xxx..
.......
`
	un := ParseUniverse(inputUniverse)
	outUniverse := un.NextStep()

	assert.Equal(t, wantUniverse, outUniverse.ToString())
}

func TestEvolveExample(t *testing.T) {
	wantUniverse := `....
....
xxx.
....
....
`

	outUniverse1 := exampleUniverse.Evolve(3)
	outUniverse2 := exampleUniverse.Evolve(4)

	assert.Equal(t, wantUniverse, outUniverse1.ToString())
	assert.Equal(t, exampleUniverse.ToString(), outUniverse2.ToString())
}

func TestEvolveRandom(t *testing.T) {
	inputUniverse := `.....x...
.........
..x...x..
.xx......
.....xx..
.x..x....
....x...x
.........
x....x...`

	wantUniverse := `.........
.........
.xx......
x..x.xx..
.xxxxxx..
.........
.........
.........
.........
`

	un := ParseUniverse(inputUniverse)
	outUniverse := un.Evolve(2)

	assert.Equal(t, wantUniverse, outUniverse.ToString())
}

func TestEvolveQuad(t *testing.T) {
	inputUniverse := `........
.xx..xx.
.x..x.x.
..x.....
.....x..
.x.x..x.
.xx..xx.
........`

	wantUniverse := `........
.xx..xx.
.x.x..x.
.....x..
..x.....
.x..x.x.
.xx..xx.
........
`

	un := ParseUniverse(inputUniverse)
	firstIteration := un.Evolve(1)
	secondIteration := un.Evolve(2)

	assert.Equal(t, wantUniverse, firstIteration.ToString())
	assert.Equal(t, inputUniverse+"\n", secondIteration.ToString())
}

func TestEvolveGlider(t *testing.T) {
	inputUniverse := `........
..x.....
...x....
.xxx....
........
........
........
........`

	wantUniverse1 := `........
........
........
....x...
.....x..
...xxx..
........
........
`
	wantUniverse2 := `........
........
........
........
........
......x.
.......x
.....xxx
`
	wantUniverse3 := `........
........
........
........
.....x..
......x.
....xxx.
........
`

	un := ParseUniverse(inputUniverse)
	u1 := un.Evolve(8)
	u2 := un.Evolve(16)
	u3 := un.Evolve(556)

	assert.Equal(t, wantUniverse1, u1.ToString())
	assert.Equal(t, wantUniverse2, u2.ToString())
	assert.Equal(t, wantUniverse3, u3.ToString())
}
