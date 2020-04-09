package main

import (
	"flag"
	"fmt"

	"github.com/ajagnic/zipf/chart"
	"github.com/ajagnic/zipf/file"
)

var filepath string

func init() {
	flag.StringVar(&filepath, "f", "", "Path for file to process")
}

func main() {
	flag.Parse()
	rmap := file.Process(filepath)
	fmt.Println(rmap)
	chart.Render(rmap)
}
