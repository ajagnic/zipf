package main

import (
	"os"

	"github.com/ajagnic/zipf/file"
)

func main() {
	filepaths := os.Args[1:]
	for _, path := range filepaths {
		file.Process(path)
	}
}
