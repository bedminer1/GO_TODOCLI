package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type statsFunc func(data []float64) float64

func sum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func csv2float(r io.Reader, column int) ([]float64, error) {
	cr := csv.NewReader(r)
	cr.ReuseRecord = true
	column-- // assume 1-indexing

	var data []float64
	for i := 0; ;i++ {
		row, err := cr.Read()
		if err == io.EOF { // variable limit of loop
			break
		}
		if err != nil {
			return nil, fmt.Errorf("cannot read data from file: %w", err)
		}

		if i == 0 { // skip header
			continue
		}

		if len(row) <= column {
			return nil, fmt.Errorf("%w: File only has %d columns", ErrInvalidColumn, len(row))
		}

		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}

		data = append(data, v)
	}

	return data, nil
}
