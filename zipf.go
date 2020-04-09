package main

import (
	"fmt"
	"os"

	"github.com/ajagnic/zipf/file"
)

func main() {
	argslice := os.Args[1:]
	for _, arg := range argslice {
		wordmap := file.Scan(arg)
		fmt.Println(wordmap)
		file.Normalize(wordmap)
	}
}
