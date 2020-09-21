package gnumeric_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus/gnumeric"
)

func TestBasicRead(t *testing.T) {
	f, err := os.Open("testdata/sample.gnumeric")
	assert.NoError(t, err)
	defer f.Close()
	workbook, err := gnumeric.Read(f)
	assert.NoError(t, err)
	assert.NotNil(t, workbook)
}

func BenchmarkBasicRead(b *testing.B) {
	f, _ := os.Open("testdata/sample.gnumeric")
	defer f.Close()
	for n := 0; n < b.N; n++ {
		gnumeric.Read(f)
	}
}
