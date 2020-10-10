package csv_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus"
	"github.com/wregis/calculus/csv"
)

func TestReadSimpleString(t *testing.T) {
	const input = "\"First\",\"Second\",\"Third\",\"Fourth\"\n\n\nFoo,Bar,Baz,?\r\n1,2,3,4,,5"
	var golden = map[string]string{
		"A1": "First", "B1": "Second", "C1": "Third", "D1": "Fourth",
		"A4": "Foo", "B4": "Bar", "C4": "Baz", "D4": "?",
		"A5": "1", "B5": "2", "C5": "3", "D5": "4", "F5": "5",
	}
	workbook, err := csv.Read(strings.NewReader(input))
	assert.NoError(t, err)
	assert.NotNil(t, workbook)
	for key := range golden {
		t.Run(key, func(t *testing.T) {
			assert.Equal(t, golden[key], workbook.ActiveSheet().ValueByRef(key))
		})
	}
}

func TestCustomRead(t *testing.T) {
	const input = "💬 Nothing to do here\n1,2,3,4,5\n💬 Neither here\n5,4,3,2,1\n%All done%"
	var golden = map[string]string{
		"A1": "1", "B1": "2", "C1": "3", "D1": "4", "E1": "5",
		"A2": "5", "B2": "4", "C2": "3", "D2": "2", "E2": "1",
		"A3": "All done",
	}
	csv := csv.File{
		Delimiter: ",",
		Enclosure: "%",
		Comment:   "💬",
	}
	workbook, err := csv.Read(strings.NewReader(input))
	assert.NoError(t, err)
	assert.NotNil(t, workbook)
	for key := range golden {
		t.Run(key, func(t *testing.T) {
			assert.Equal(t, golden[key], workbook.ActiveSheet().ValueByRef(key))
		})
	}
}

func TestHints(t *testing.T) {
	const input = "1,2.3,true,Hello\n42,-0.0002,false,World\n-99,0.12345,1,\"Foo\""
	var golden = map[string]interface{}{
		"A1": int64(1), "B1": 2.3, "C1": true, "D1": "Hello",
		"A2": int64(42), "B2": -0.0002, "C2": false, "D2": "World",
		"A3": int64(-99), "B3": 0.12345, "C3": true, "D3": "Foo",
	}
	config := csv.New()
	config.Hints = []csv.Hint{
		csv.Hint{Type: calculus.CellValueTypeInteger},
		csv.Hint{Type: calculus.CellValueTypeFloat},
		csv.Hint{Type: calculus.CellValueTypeBoolean},
	}
	workbook, err := config.Read(strings.NewReader(input))
	assert.NoError(t, err)
	assert.NotNil(t, workbook)
	for key := range golden {
		t.Run(key, func(t *testing.T) {
			assert.Equal(t, golden[key], workbook.ActiveSheet().ValueByRef(key))
		})
	}
}

func TestReadFile(t *testing.T) {
	file, _ := os.Open("testdata/countries.csv")
	workbook, err := csv.Read(file)
	assert.NoError(t, err)
	assert.NotNil(t, workbook)
}

func BenchmarkReadBasicString(b *testing.B) {
	const input = "\"First\",\"Second\",\"Third\",\"Fourth\"\n\n\nFoo,Bar,Baz,?\r\n1,2,3,4,,5"
	for n := 0; n < b.N; n++ {
		csv.Read(strings.NewReader(input))
	}
}

func BenchmarkReadLargeFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		file, _ := os.Open("testdata/countries.csv")
		csv.Read(file)
	}
}
