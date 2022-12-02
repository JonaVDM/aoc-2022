package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(file string) []string {
	f, err := os.ReadFile(fmt.Sprintf("inputs/%s", file))
	path, err1 := os.Getwd()
	fmt.Println(path, err1)
	if err != nil {
		panic("could not read input file")
	}

	inp := string(f)
	inp = strings.TrimSpace(inp)

	return strings.Split(inp, "\n")
}
