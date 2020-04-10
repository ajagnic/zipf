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

func format(ratiomap map[float64][]string) (wordsdesc []string, ratiosdesc []float64) {
	ratioslc := []float64{}
	for k, _ := range ratiomap {
		ratioslc = append(ratioslc, k)
	}
	sort.Float64s(ratioslc)
	wordsdesc = []string{}
	ratiosdesc = []float64{}
	for i := len(ratioslc) - 1; i >= 0; i-- {
		wordsdesc = append(wordsdesc, ratiomap[ratioslc[i]][0])
		ratiosdesc = append(ratiosdesc, ratioslc[i])
	}
	return
}
