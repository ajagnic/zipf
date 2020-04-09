package file

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ajagnic/zipf/check"
)

func Scan(path string) (wordmap map[string]int) {
	filep, err := os.Open(path)
	defer filep.Close()
	check.Fatal(err)
	scannerp := bufio.NewScanner(filep)
	scannerp.Split(bufio.ScanWords)
	wordmap = make(map[string]int)
	for scannerp.Scan() {
		word := scannerp.Text()
		wordmap[word]++
	}
	err = scannerp.Err()
	check.Fatal(err)
	return
}

func Normalize(wordmap map[string]int) {
	for k := range wordmap {
		fmt.Println(k)
	}
}
