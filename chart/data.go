package chart

import (
	"fmt"
	"sort"
)

func Render(ratiomap map[float64][]string) {
	format(ratiomap)
}

func format(ratiomap map[float64][]string) {
	ratioslice := []float64{}
	for k, _ := range ratiomap {
		ratioslice = append(ratioslice, k)
	}
	sort.Float64s(ratioslice)
	sortedwords := []string{}
	for i := len(ratioslice) - 1; i >= 0; i-- {
		for _, word := range ratiomap[ratioslice[i]] {
			sortedwords = append(sortedwords, word)
		}
	}
	fmt.Println(sortedwords)
}
