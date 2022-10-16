package data

import (
	"github.com/samurai-of-honor/game-of-life/internal/core"
	"io"
	"io/ioutil"
	"os"
)

func ReadFile(inputFile string) string {
	file, _ := os.Open(inputFile)
	defer file.Close()

	buf, _ := io.ReadAll(file)
	return string(buf)
}

func WriteFile(finalUniverse core.Universe) {
	_ = os.Mkdir("out", 0777)
	_ = ioutil.WriteFile("./out/finalUniverse.txt", []byte(finalUniverse.ToString()), 0777)
}
