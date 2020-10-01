package ods_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus/ods"
)

func TestReadSample(t *testing.T) {
	workbook, err := ods.ReadFile("testdata/sample.ods")
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
	assert.Equal(t, 1.0, sheet.ValueByRef("C3"))
}
