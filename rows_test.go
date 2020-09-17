package calculus_test

import (
	"testing"

	"github.com/wregis/calculus"
)

func BenchmarkRows_SetCell_NewRow(b *testing.B) {
	cell := calculus.NewCell("Hello")
	rows := calculus.NewRows()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		rows.SetCell(n, n, cell)
	}
}
