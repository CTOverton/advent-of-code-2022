package utils

import (
	"log"
	"os"
)

func ReadInputFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
