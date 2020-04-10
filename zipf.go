package main

import (
	"flag"

	"github.com/ajagnic/zipf/chart"
	"github.com/ajagnic/zipf/file"
)

var filepath string

func init() {
	flag.StringVar(&filepath, "f", "", "Path for file to process")
}

func main() {
	flag.Parse()
	ratiomap := file.Process(filepath)
	chart.Render(ratiomap)
}
