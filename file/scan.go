package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ajagnic/zipf/check"
)

func Process(path string) {
	wordmap := scan(path)
	normalize(wordmap)
	fmt.Println(wordmap)
}

func scan(path string) (wordmap map[string]int) {
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

func normalize(wordmap map[string]int) {
	for key, val := range wordmap {
		lowerkey := strings.ToLower(key)
		newkey := strings.Trim(lowerkey, ";:,.?!()")
		if newkey != key {
			wordmap[newkey] += val
			delete(wordmap, key)
		}
	}
}
