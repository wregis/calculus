package calculus_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus"
)

func TestParseCoordinate(t *testing.T) {
	cases := []struct {
		input  string
		output [2]int
	}{
		{"A1", [2]int{0, 0}},
		{"Y17", [2]int{16, 24}},
		{"AH12", [2]int{11, 33}},
		{"BD27", [2]int{26, 55}},
		{"ALL1000", [2]int{999, 999}},
	}
	for _, testCase := range cases {
		t.Run(testCase.input, func(t *testing.T) {
			row, column, err := calculus.ParseCoordinate(testCase.input)
			assert.NoError(t, err)
			assert.Equal(t, testCase.output[0], row)
			assert.Equal(t, testCase.output[1], column)
		})
	}
}

func TestCoordinate(t *testing.T) {
	cases := []struct {
		output string
		input  [2]int
	}{
		{"A1", [2]int{0, 0}},
		{"Y17", [2]int{16, 24}},
		{"AH12", [2]int{11, 33}},
		{"BD27", [2]int{26, 55}},
		{"ALL1000", [2]int{999, 999}},
	}
	for _, testCase := range cases {
		t.Run(testCase.output, func(t *testing.T) {
			key, err := calculus.Coordinate(testCase.input[0], testCase.input[1])
			assert.NoError(t, err)
			assert.Equal(t, testCase.output, key)
		})
	}
}
