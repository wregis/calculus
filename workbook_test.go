package calculus_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus"
)

func TestWorkbookNew(t *testing.T) {
	workbook := calculus.New()

	assert.NotNil(t, workbook)
	assert.NotNil(t, workbook.Properties())
	assert.Nil(t, workbook.ActiveSheet())
	assert.Empty(t, workbook.Sheets())
}

func TestWorkbookSheetHandling(t *testing.T) {
	workbook := calculus.New()

	assert.Error(t, workbook.RemoveSheet("Hello"))

	sheet1, err := workbook.AddSheet("Foo")
	assert.NoError(t, err)
	assert.NotNil(t, sheet1)
	assert.Len(t, workbook.Sheets(), 1)

	sheet2, err := workbook.AddSheet("Foo")
	assert.Error(t, err)
	assert.Nil(t, sheet2)
	assert.Len(t, workbook.Sheets(), 1)

	assert.Same(t, sheet1, workbook.ActiveSheet())

	sheet3, err := workbook.AddSheetFirst("Bar")
	assert.NoError(t, err)
	assert.NotNil(t, sheet3)
	assert.Len(t, workbook.Sheets(), 2)

	assert.Same(t, sheet1, workbook.Sheet("Foo"))
	assert.Same(t, sheet1, workbook.SetActive("Foo"))

	assert.NoError(t, workbook.RemoveSheet("Foo"))
	assert.Error(t, workbook.RemoveSheet("Foo"))
	assert.Len(t, workbook.Sheets(), 1)

	assert.NoError(t, workbook.RemoveSheet("Bar"))
	assert.Empty(t, workbook.Sheets())
}

func BenchmarkWorkbookBasicComposition(b *testing.B) {
	for n := 0; n < b.N; n++ {
		workbook := calculus.New()
		workbook.Properties().SetApplication("Calculus Test")
		workbook.Properties().SetCreated(time.Now())
		workbook.Properties().SetDescription("Calculus test unit")

		sheet, _ := workbook.AddSheet("TestSheet1")
		sheet.SetValue(0, 0, "Hello")
		sheet.SetValue(0, 1, "World")
		sheet.SetValue(1, 0, "Foo")
		sheet.SetValue(1, 1, "Bar")
		sheet.SetValue(1, 2, "Baz")
		sheet.SetValue(2, 0, 3.1415926535)
		sheet.SetValue(2, 1, 42)
		sheet.SetValue(2, 2, true)
	}
}
