package renderer

import (
	"fmt"

	"github.com/fikriauliya/yamlexplorer/entity"
	"github.com/montanaflynn/stats"
)

func transpose(a [][]string) [][]string {
	res := make([][]string, len(a[0]))
	if len(res) == 0 {
		return [][]string{{}}
	}
	for i := 0; i < len(res); i++ {
		res[i] = make([]string, len(a))
	}
	for i, row := range a {
		for j, col := range row {
			res[j][i] = col
		}
	}
	return res
}

func lenMatrix(a [][]string) [][]float64 {
	res := make([][]float64, len(a))

	for i := 0; i < len(a); i++ {
		res[i] = make([]float64, len(a[i]))
		for j := 0; j < len(a[i]); j++ {
			res[i][j] = float64(len(a[i][j]))
		}
	}
	return res
}

func medianVector(a [][]float64) (*[]float64, error) {
	res := make([]float64, len(a))
	for i, row := range a {
		med, err := stats.Median(row)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate median: %s", err)
		}
		res[i] = med
	}
	return &res, nil
}

func max(v []float64) float64 {
	max := v[0]
	for _, item := range v {
		if item > max {
			max = item
		}
	}
	return max
}

func calculateWidths(m [][]string, maxWidth int) ([]int, error) {
	columns := lenMatrix(transpose(m))
	medianLengths, err := medianVector(columns)
	if err != nil {
		return []int{}, err
	}

	widths := make([]int, len(columns))
	usedWidth := 0
	for i := range columns {
		remWidth := maxWidth - usedWidth
		width := 0
		if i == len(columns)-1 {
			width = remWidth
		} else {
			medianLen := int((*medianLengths)[i])
			maxLen := int(max(columns[i]))

			width = medianLen
			if maxLen-width <= 2 {
				width = maxLen
			}

			if remWidth < width {
				width = remWidth
			}
			usedWidth += width
		}
		widths[i] = width
	}
	return widths, nil
}

func Resize(t *entity.Table, maxWidth int) ([]int, error) {
	widths, err := calculateWidths(t.Body, maxWidth)
	if err != nil {
		return []int{}, err
	}
	return widths, nil
}
