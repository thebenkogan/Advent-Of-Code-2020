package input

import (
	"fmt"
	"os"
)

func GetInput() string {
	day := os.Args[1]
	var inputFileName string
	if len(os.Args) > 2 && os.Args[2] == "test" {
		inputFileName = "test"
	} else {
		inputFileName = "in"
	}

	pwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/cmd/%s/%s.txt", pwd, day, inputFileName)
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
