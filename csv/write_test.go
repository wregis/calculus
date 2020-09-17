package csv_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus"
	"github.com/wregis/calculus/csv"
)

func TestEmptyComposition(t *testing.T) {
	workbook := calculus.New()

	buf := new(bytes.Buffer)
	err := csv.Write(workbook, buf)
	assert.Error(t, err)
	assert.Empty(t, buf.String())
}

func TestSimpleStringComposition(t *testing.T) {
	const golden = "Hello,World\nFoo,Bar,Baz"

	workbook := calculus.New()
	sheet, _ := workbook.AddSheet("TestSheet1")
	sheet.SetValue(0, 0, "Hello")
	sheet.SetValue(0, 1, "World")
	sheet.SetValueByRef("A2", "Foo")
	sheet.SetValueByRef("B2", "Bar")
	sheet.SetValueByRef("C2", "Baz")

	buf := new(bytes.Buffer)
	err := csv.Write(workbook, buf)
	assert.NoError(t, err)
	assert.Equal(t, golden, buf.String())
}

func TestCustomParamsComposition(t *testing.T) {
	const golden = "\"Hi there!\"\n;\"Go ➡️\"\n;;\"See a 🌈\"\n;;;\"No escape \\\""

	workbook := calculus.New()
	sheet, _ := workbook.AddSheet("TestSheet1")
	sheet.SetValue(0, 0, "Hi there!")
	sheet.SetValue(1, 1, "Go ➡️")
	sheet.SetValue(2, 2, "See a 🌈")
	sheet.SetValue(3, 3, "No escape \\")

	buf := new(bytes.Buffer)
	csv := csv.File{
		Delimiter:  ";",
		Enclosure:  "\"",
		EscapeChar: '\\',
	}
	err := csv.Write(workbook, buf)
	assert.NoError(t, err)
	assert.Equal(t, golden, buf.String())
}

func TestVariousTypesComposition(t *testing.T) {
	const golden = "One,Two,Three\n3.1415926535\n\ntrue,false"

	workbook := calculus.New()
	sheet, _ := workbook.AddSheet("TestSheet1")
	sheet.SetValue(0, 0, "Zero")
	sheet.SetValue(0, 0, "One")
	sheet.SetValue(0, 1, "Two")
	sheet.SetValue(0, 2, "Three")
	sheet.SetValue(1, 0, true)
	sheet.SetValue(1, 0, 2)
	sheet.SetValue(1, 0, 3.1415926535)
	sheet.SetValue(3, 0, true)
	sheet.SetValue(3, 1, false)

	buf := new(bytes.Buffer)
	err := csv.Write(workbook, buf)
	assert.NoError(t, err)
	assert.Equal(t, golden, buf.String())
}
