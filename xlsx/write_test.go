package xlsx_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus"
	"github.com/wregis/calculus/xlsx"
)

func TestEmptyComposition(t *testing.T) {
	workbook := calculus.New()

	buf := new(bytes.Buffer)
	err := xlsx.Write(workbook, buf)
	assert.Error(t, err)
	assert.Empty(t, buf.Bytes())
}

func TestSimpleStringComposition(t *testing.T) {
	workbook := calculus.New()
	workbook.Properties().SetApplication("Calculus Test")
	workbook.Properties().SetCreated(time.Now())
	workbook.Properties().SetDescription("Calculus test unit")
	sheet, _ := workbook.AddSheet("TestSheet1")
	sheet.SetValue(0, 0, "Hello")
	sheet.SetValue(0, 1, "World")
	sheet.SetValueByRef("A2", "Foo")
	sheet.SetValueByRef("B2", "Bar")
	sheet.SetValueByRef("C2", "Baz")
	sheet.SetValue(2, 0, 3.1415926535)
	sheet.SetValue(2, 1, 42)
	sheet.SetValue(2, 2, true)

	buf := new(bytes.Buffer)
	err := xlsx.Write(workbook, buf)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf.Bytes())
}

func BenchmarkWrite(b *testing.B) {
	workbook := calculus.New()
	sheet, _ := workbook.AddSheet("TestSheet1")
	sheet.SetValue(0, 0, "Hello")
	sheet.SetValue(0, 1, "World")
	sheet.SetValueByRef("A2", "Foo")
	sheet.SetValueByRef("B2", "Bar")
	sheet.SetValueByRef("C2", "Baz")
	sheet.SetValue(2, 0, 3.1415926535)
	sheet.SetValue(2, 1, 42)
	sheet.SetValue(2, 2, true)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		xlsx.Write(workbook, buf)
	}
}
