package csv_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus/csv"
)

func TestReadSimpleString(t *testing.T) {
	const input string = "\"First\",\"Second\",\"Third\",\"Fourth\"\n\n\nFoo,Bar,Baz,?\r\n1,2,3,4,,5"
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
	const input string = "💬 Nothing to do here\n1,2,3,4,5\n💬 Neither here\n5,4,3,2,1\n%All done%"
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
