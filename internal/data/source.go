package data

import (
	"io"
	"io/ioutil"
	"ip-0x-lab-1-samurai-of-honor/internal/core"
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
