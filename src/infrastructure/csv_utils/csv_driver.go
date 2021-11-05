package csv_utils

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

// ReadCSV expects the path of a CSV file and writes new lines to the file
func AppendCSV(path string, lines [][]string) error {
	csvFile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(csvFile)
	defer csvFile.Close()
	defer writer.Flush()

	err = writer.WriteAll(lines)
	if err != nil {
		return err
	}

	return nil
}

// CountCSVLines expects a CSV file in filename and count the number of lines
func CountCSVLines(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return -1, err
	}

	return len(lines), nil
}
