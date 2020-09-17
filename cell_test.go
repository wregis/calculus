package calculus_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus"
)

func TestCell(t *testing.T) {
	testCases := []struct {
		t calculus.CellValueType
		v interface{}
	}{
		{calculus.CellValueTypeString, "Hello"},
		{calculus.CellValueTypeFormula, "=SUM(1, 2)"},
		{calculus.CellValueTypeNumber, 1.41421356237},
		{calculus.CellValueTypeNumber, 128},
		{calculus.CellValueTypeBoolean, false},
	}
	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("%d-%T", index, testCase.t), func(t *testing.T) {
			cell := calculus.NewCell(testCase.v)
			assert.Equal(t, cell.Type(), testCase.t)
			assert.Equal(t, cell.Value(), testCase.v)
		})
	}
}

func BenchmarkNewCellString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		calculus.NewCell("Hello")
	}
}

func BenchmarkNewCellFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		calculus.NewCell(2.71828182846)
	}
}

func BenchmarkNewCellInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		calculus.NewCell(42)
	}
}

func BenchmarkNewCellBoolean(b *testing.B) {
	for n := 0; n < b.N; n++ {
		calculus.NewCell(true)
	}
}
