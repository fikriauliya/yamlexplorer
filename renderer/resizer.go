package renderer

import (
	"fmt"

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

func Resize(body [][]string, maxWidth int) (*[][]string, error) {
	columns := lenMatrix(transpose(body))
	medianLengths, err := medianVector(columns)
	if err != nil {
		return nil, err
	}

	res := make([][]string, len(body))
	for i, row := range body {
		res[i] = make([]string, len(row))
		for j := range row {
			minLen := int((*medianLengths)[j])
			if minLen > len(body[i][j]) {
				minLen = len(body[i][j])
			}
			res[i][j] = body[i][j][:minLen]
		}
	}
	return &res, nil
}
