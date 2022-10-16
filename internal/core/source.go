package core

import (
	"io"
	"io/ioutil"
	"math/rand"
	"os"
)

func GenerateRandomUniverse(width, height int) Universe {
	universe := make(Universe, height)

	for i := range universe {
		universe[i] = make([]bool, width)

		for j := range universe[i] {
			universe[i][j] = rand.Intn(5) == 0
		}
	}

	return universe
}

func ReadFile(inputFile string) string {
	file, _ := os.Open(inputFile)
	defer file.Close()

	buf, _ := io.ReadAll(file)
	return string(buf)
}

func WriteFile(finalUniverse Universe) {
	_ = os.Mkdir("out", 0777)
	_ = ioutil.WriteFile("./out/finalUniverse.txt", []byte(finalUniverse.ToString()), 0777)
}
