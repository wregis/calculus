package xlsx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus/xlsx"
)

func TestReadSample(t *testing.T) {
	workbook, err := xlsx.ReadFile("testdata/sample.xlsx")
	assert.NoError(t, err)
	assert.NotNil(t, workbook)

	assert.NotEmpty(t, workbook.Sheets())
	assert.Len(t, workbook.Sheets(), 1)
	assert.NotNil(t, workbook.ActiveSheet())
	assert.NotNil(t, workbook.Sheet("TestSheet1"))
	assert.Same(t, workbook.ActiveSheet(), workbook.Sheet("TestSheet1"))

	sheet := workbook.Sheet("TestSheet1")
	assert.Equal(t, "Hello", sheet.ValueByRef("A1"))
	assert.Equal(t, "World", sheet.ValueByRef("B1"))
	assert.Equal(t, "Foo", sheet.ValueByRef("A2"))
	assert.Equal(t, "Bar", sheet.ValueByRef("B2"))
	assert.Equal(t, "Baz", sheet.ValueByRef("C2"))
	assert.Equal(t, 3.1415926535, sheet.ValueByRef("A3"))
	assert.Equal(t, 42.0, sheet.ValueByRef("B3"))
	assert.Equal(t, true, sheet.ValueByRef("C3"))
}

func TestReadLocations(t *testing.T) {
	// https://population.un.org/wpp/Download/Files/4_Metadata/WPP2019_F01_LOCATIONS.XLSX
	// 11184 cells (values, formulas and/or styles)
	workbook, err := xlsx.ReadFile("testdata/un_locations.xlsx")
	assert.NoError(t, err)
	assert.NotNil(t, workbook)

	locSheet := workbook.Sheet("Location")
	assert.NotNil(t, locSheet)
	assert.Equal(t, "Geographic region", locSheet.ValueByRef("G35"))
	assert.Equal(t, "Brazil", locSheet.ValueByRef("B207"))
	assert.Equal(t, "CAN", locSheet.ValueByRef("E303"))
	assert.Equal(t, "Country/Area", locSheet.ValueByRef("G294"))

	dbSheet := workbook.Sheet("DB")
	assert.NotNil(t, dbSheet)
	assert.Equal(t, "Africa", dbSheet.ValueByRef("B19"))
	assert.Equal(t, "Angola", dbSheet.ValueByRef("B49"))
	assert.Equal(t, "Western Asia", dbSheet.ValueByRef("K108"))
	assert.Equal(t, "Oceania", dbSheet.ValueByRef("Q226"))

	notesSheet := workbook.Sheet("NOTES")
	assert.NotNil(t, notesSheet)
	assert.Equal(t, "p", notesSheet.ValueByRef("A18"))
	assert.Equal(t, 1.0, notesSheet.ValueByRef("A19"))
	assert.Contains(t, notesSheet.ValueByRef("B48"), "Vatican")
}

func BenchmarkReadSample(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xlsx.ReadFile("testdata/sample.xlsx")
	}
}

func BenchmarkReadLocations(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xlsx.ReadFile("testdata/un_locations.xlsx")
	}
}
