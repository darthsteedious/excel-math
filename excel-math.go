package excelmath

import (
	"errors"
	"math"
)

// GeoMean - Calculates the geometric mean of a series
func GeoMean(series []float64) (float64, error) {
	n := len(series)
	if n == 0 {
		return 0, errors.New("")
	}

	acc := 1.0
	for _, v := range series {
		acc = acc * v
	}

	return math.Pow(acc, 1/float64(n)), nil
}
