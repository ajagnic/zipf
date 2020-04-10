package file

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

func Process(path string) (ratiomap map[float64][]string) {
	wordmap := scan(path)
	normalize(wordmap)
	ratiomap = sort(wordmap)
	return
}

func scan(path string) (wordmap map[string]int) {
	filep, err := os.Open(path)
	defer filep.Close()
	if err != nil {
		log.Fatalf("\nFatalError: %v", err)
	}
	scannerp := bufio.NewScanner(filep)
	scannerp.Split(bufio.ScanWords)
	wordmap = make(map[string]int)
	for scannerp.Scan() {
		word := scannerp.Text()
		wordmap[word]++
	}
	err = scannerp.Err()
	if err != nil {
		log.Fatalf("\nFatalError: %v", err)
	}
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
	total := 0.0
	for _, val := range wordmap {
		total += float64(val)
	}
	ratiomap = make(map[float64][]string)
	for key, val := range wordmap {
		ratio := (float64(val) / total) * 100
		decimal := math.Round(ratio*100) / 100
		ratiomap[decimal] = append(ratiomap[decimal], key)
	}
	return
}
