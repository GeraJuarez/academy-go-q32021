package common

import (
	"log"
	"os"
	"reflect"
	"testing"
)

const CSV_TEST_PATH = "./test.csv"

var cases = []struct {
	vals []string
	err  error
}{
	{[]string{"1", "test01"}, nil},
	{[]string{"2", "test02"}, nil},
}

func TestMain(m *testing.M) {
	code := m.Run()
	err := os.Remove(CSV_TEST_PATH)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

func TestAppendCSV(t *testing.T) {
	for _, c := range cases {
		err := AppendCSV(CSV_TEST_PATH, [][]string{c.vals})
		if err != c.err {
			t.Errorf("Expected %v,got %v", c.err, err)
		}
	}

}

func TestReadCSV(t *testing.T) {
	lines, err := ReadCSV(CSV_TEST_PATH)
	if err != nil {
		t.Fatal(err)
	}
	for idx, c := range cases {
		got := lines[idx]
		if !reflect.DeepEqual(got, c.vals) {
			t.Errorf("Expected %s,got %s", c.vals, got)
		}
	}
}
