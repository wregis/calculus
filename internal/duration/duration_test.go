package duration_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus/internal/duration"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		source string
		result time.Duration
	}{
		{"P0D", 0},
		{"PT0S", 0},
		{"PT1S", time.Second},
		{"PT1M", time.Minute},
		{"PT1H", time.Hour},
		{"P1D", time.Hour * 24},
		{"P1W", time.Hour * 24 * 7},
		{"P1M", time.Hour * 24 * 30},
		{"P1Y", time.Hour * 24 * 365},
		{"PT654H30M21S", (654 * time.Hour) + (30 * time.Minute) + (21 * time.Second)},
		{"PT13H22M11S", (13 * time.Hour) + (22 * time.Minute) + (11 * time.Second)},
		{"P10Y5M8DT5H10M6S", (10 * 24 * 365 * time.Hour) + (5 * 24 * 30 * time.Hour) + (8 * 24 * time.Hour) + (5 * time.Hour) + (10 * time.Minute) + (6 * time.Second)},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%s", testCase.source), func(t *testing.T) {
			d, err := duration.Parse(testCase.source)
			assert.NoError(t, err)
			assert.Equal(t, testCase.result, d)
		})
	}
}

func BenchmarkParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		duration.Parse("P10Y5M8DT5H10M6S")
	}
}
