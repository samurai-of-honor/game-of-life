package data

import (
	"github.com/samurai-of-honor/game-of-life/internal/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	inputStr := `3
5 4
.....
.x...
.x.x.
.....`

	n, universe := ParseInput(inputStr)

	assert.Equal(t, 3, n)
	assert.False(t, universe[0][0])
	assert.True(t, universe[1][1])
}

func TestParseInput2(t *testing.T) {
	inputStr := `159
 
......x.
.x...xxx
.x.x...x
.....x..
.x..x...`

	n, universe := ParseInput(inputStr)

	assert.Equal(t, 159, n)
	assert.Equal(t, 5, len(universe))

	assert.Equal(t, []bool{false, true, false, true, false, false, false, true}, universe[2])
}

func TestUniverseToString(t *testing.T) {
	testUniverse := core.Universe{
		{false, false, false},
		{false, true, false},
		{false, false, true},
	}
	wantStr := `...
.x.
..x
`

	str := testUniverse.ToString()

	assert.Equal(t, wantStr, str)
}

func TestParseUniverse(t *testing.T) {
	inputStr := `....
.x..
.x.x.
.....`

	universe := core.ParseUniverse(inputStr)

	assert.Equal(t, 4, len(universe))
	assert.False(t, universe[3][0])
	assert.True(t, universe[2][3])
}
