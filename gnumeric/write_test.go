package gnumeric_test

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus"
	"github.com/wregis/calculus/gnumeric"
)

func TestEmptyComposition(t *testing.T) {
	workbook := calculus.New()

	buf := new(bytes.Buffer)
	err := gnumeric.Write(workbook, buf)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf.Bytes())
}

func TestVariousTypesComposition(t *testing.T) {
	workbook := calculus.New()
	sheet, _ := workbook.AddSheet("TestSheet1")
	sheet.SetValue(0, 0, "One")
	sheet.SetValue(0, 1, "Two")
	sheet.SetValue(0, 2, "Three")
	sheet.SetValue(1, 0, 3.1415926535)
	sheet.SetValue(3, 0, true)
	sheet.SetValue(3, 1, false)
	sheet.SetValue(4, 0, time.Date(2019, time.August, 13, 12, 13, 14, 0, time.UTC))

	buf := new(bytes.Buffer)
	err := gnumeric.Write(workbook, buf)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf.Bytes())

	reader, err := gzip.NewReader(bytes.NewBuffer(buf.Bytes()))
	assert.NoError(t, err)
	data, err := ioutil.ReadAll(reader)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	assert.Contains(t, string(data), "Name>TestSheet1</Name>")
	assert.Contains(t, string(data), ">One</Cell>")
	assert.Contains(t, string(data), ">Two</Cell>")
	assert.Contains(t, string(data), ">3.1415926535</Cell>")
	assert.Contains(t, string(data), ">TRUE</Cell>")
	assert.Contains(t, string(data), ">43690.50918981482</Cell>")
}

func BenchmarkComposition(b *testing.B) {
	workbook := calculus.New()
	sheet, _ := workbook.AddSheet("TestSheet1")
	sheet.SetValue(0, 0, "One")
	sheet.SetValue(0, 1, "Two")
	sheet.SetValue(0, 2, "Three")
	sheet.SetValue(1, 0, 3.1415926535)
	sheet.SetValue(3, 0, true)
	sheet.SetValue(3, 1, false)

	buf := new(bytes.Buffer)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buf.Reset()
		gnumeric.Write(workbook, buf)
	}
}
