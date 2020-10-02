package format_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wregis/calculus/internal/format"
)

func TestToTime1900(t *testing.T) {
	testCases := []struct {
		t        time.Time
		date1904 bool
		f        float64
	}{
		{time.Date(2029, time.January, 1, 0, 0, 0, 0, time.UTC), false, 47119},
		{time.Date(1930, time.January, 1, 0, 0, 0, 0, time.UTC), false, 10959},
		{time.Date(1940, time.January, 1, 0, 0, 0, 0, time.UTC), false, 14611},
		{time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC), false, 42005},
		{time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC), false, 36892},

		{time.Date(2018, time.January, 1, 1, 0, 0, 0, time.UTC), false, 43101.041666666664},
		{time.Date(2017, time.January, 1, 14, 0, 30, 0, time.UTC), false, 42736.58368055556},
		{time.Date(2017, time.January, 1, 14, 0, 0, 0, time.UTC), false, 42736.583333333336},

		{time.Date(2020, time.September, 22, 9, 0, 0, 0, time.UTC), false, 44096.375},
		{time.Date(2020, time.September, 22, 1, 0, 0, 0, time.UTC), false, 44096.041666666664},
		{time.Date(2020, time.September, 22, 13, 0, 0, 0, time.UTC), false, 44096.541666666664},
		{time.Date(2020, time.September, 22, 12, 30, 45, 0, time.UTC), false, 44096.52135416667},
		{time.Date(2020, time.September, 22, 11, 15, 30, 0, time.UTC), false, 44096.46909722222},
		{time.Date(2020, time.September, 22, 16, 45, 0, 0, time.UTC), false, 44096.697916666664},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%s %v", testCase.t.String(), testCase.date1904), func(t *testing.T) {
			f := format.ToTime1900(testCase.t, testCase.date1904)
			assert.Equal(t, testCase.f, f)
		})
	}
}

func TestFromTime1900(t *testing.T) {
	testCases := []struct {
		f        float64
		date1904 bool
		t        time.Time
	}{
		{47119, false, time.Date(2029, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{10959, false, time.Date(1930, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{14611, false, time.Date(1940, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{42005, false, time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{36892, false, time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)},

		{43101.041666666664, false, time.Date(2018, time.January, 1, 1, 0, 0, 0, time.UTC)},
		{42736.58368055556, false, time.Date(2017, time.January, 1, 14, 0, 30, 0, time.UTC)},
		{42736.583333333336, false, time.Date(2017, time.January, 1, 14, 0, 0, 0, time.UTC)},

		{44096.375, false, time.Date(2020, time.September, 22, 9, 0, 0, 0, time.UTC)},
		{44096.041666666664, false, time.Date(2020, time.September, 22, 1, 0, 0, 0, time.UTC)},
		{44096.541666666664, false, time.Date(2020, time.September, 22, 13, 0, 0, 0, time.UTC)},
		{44096.52135416667, false, time.Date(2020, time.September, 22, 12, 30, 45, 0, time.UTC)},
		{44096.46909722222, false, time.Date(2020, time.September, 22, 11, 15, 30, 0, time.UTC)},
		{44096.697916666664, false, time.Date(2020, time.September, 22, 16, 45, 0, 0, time.UTC)},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%s %v", testCase.t.String(), testCase.date1904), func(t *testing.T) {
			tt := format.FromTime1900(testCase.f, testCase.date1904)
			assert.Equal(t, testCase.t.Unix(), tt.Unix())
		})
	}
}

func BenchmarkToTime1900(b *testing.B) {
	cases := []time.Time{
		time.Date(2029, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1930, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1940, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC),

		time.Date(2018, time.January, 1, 1, 0, 0, 0, time.UTC),
		time.Date(2017, time.January, 1, 14, 0, 30, 0, time.UTC),
		time.Date(2017, time.January, 1, 14, 0, 0, 0, time.UTC),

		time.Date(2020, time.September, 22, 9, 0, 0, 0, time.UTC),
		time.Date(2020, time.September, 22, 1, 0, 0, 0, time.UTC),
		time.Date(2020, time.September, 22, 13, 0, 0, 0, time.UTC),
		time.Date(2020, time.September, 22, 12, 30, 45, 0, time.UTC),
		time.Date(2020, time.September, 22, 11, 15, 30, 0, time.UTC),
		time.Date(2020, time.September, 22, 16, 45, 0, 0, time.UTC),
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		format.ToTime1900(cases[n%len(cases)], false)
	}
}

func BenchmarkFromTime1900(b *testing.B) {
	cases := []float64{
		47119, 10959, 14611, 42005, 36892,
		43101.041666666664, 42736.58368055556, 42736.583333333336,
		44096.375, 44096.041666666664, 44096.541666666664, 44096.52135416667, 44096.46909722222, 44096.697916666664,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		format.FromTime1900(cases[n%len(cases)], false)
	}
}
