package calculus_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus"
)

func TestCell(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		typ        uint8
		in, result interface{}
	}{
		{calculus.CellValueTypeEmpty, nil, nil},
		{calculus.CellValueTypeBoolean, false, false},
		{calculus.CellValueTypeFloat, 1.41421356237, 1.41421356237},
		{calculus.CellValueTypeInteger, 128, 128},
		{calculus.CellValueTypeString, "Hello", "Hello"},
		{calculus.CellValueTypeString, []rune("こんにちは"), []rune("こんにちは")},
		{calculus.CellValueTypeString, []float32{}, "[]"},
		{calculus.CellValueTypeDate, now, now},
		{calculus.CellValueTypeDuration, time.Minute, time.Minute},
	}
	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("%d-%T", index, testCase.typ), func(t *testing.T) {
			cell := calculus.NewCell(testCase.in)
			assert.Equal(t, cell.Type(), testCase.typ)
			assert.Equal(t, cell.Value(), testCase.result)
		})
	}
}

func BenchmarkNewCellString(b *testing.B) {
	cases := []interface{}{
		nil,
		false,
		true,
		3.1415926535,
		time.Now(),
		time.Second,
		"مرحبا",
		[]rune("Olá"),
		[]int{2, 4, 8, 16, 32, 64, 128},
	}
	for n := 0; n < b.N; n++ {
		calculus.NewCell(cases[n%len(cases)])
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
