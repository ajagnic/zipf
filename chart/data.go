package chart

import (
	"log"
	"os"
	"sort"

	"github.com/go-echarts/go-echarts/charts"
)

func Render(ratiomap map[float64][]string) {
	words, ratios := format(ratiomap)
	bar := charts.NewBar()
	bar.AddXAxis(words).AddYAxis("% Total", ratios)
	filep, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(filep)
}

func format(ratiomap map[float64][]string) (sortedwords []string, sortedratios []float64) {
	ratioslice := []float64{}
	for k, _ := range ratiomap {
		ratioslice = append(ratioslice, k)
	}
	sort.Float64s(ratioslice)
	sortedwords = []string{}
	sortedratios = []float64{}
	for i := len(ratioslice) - 1; i >= 0; i-- {
		for _, word := range ratiomap[ratioslice[i]] {
			sortedwords = append(sortedwords, word)
			sortedratios = append(sortedratios, ratioslice[i])
		}
	}
	return
}
