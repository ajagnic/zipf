package file

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/ajagnic/zipf/check"
)

func Process(path string) {
	wordmap := scan(path)
	normalize(wordmap)
	fmt.Println(wordmap)
	ratiomap := sort(wordmap)
	fmt.Println(ratiomap)
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
		newkey := strings.Trim(strings.ToLower(key), ";:,.?!()")
		if newkey != key {
			wordmap[newkey] += val
			delete(wordmap, key)
		}
	}
}

func sort(wordmap map[string]int) (ratiomap map[float64][]string) {
	ratiomap = make(map[float64][]string)
	total := 0.0
	for _, val := range wordmap {
		total += float64(val)
	}
	for key, val := range wordmap {
		ratio := (float64(val) / total) * 100
		roundratio := math.Round(ratio*100) / 100
		ratiomap[roundratio] = append(ratiomap[roundratio], key)
	}
	return
}
