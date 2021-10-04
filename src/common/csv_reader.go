package common

import (
	"encoding/csv"
	"os"
)

// ReadCSV expects a CSV file in filename and parse it into
// a double array string
func ReadCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
