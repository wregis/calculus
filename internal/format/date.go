package format

import "time"

const secondsInADay = 24 * time.Hour

var (
	time1900Epoch         = time.Date(1899, time.December, 31, 0, 0, 0, 0, time.UTC)
	time1904Epoch         = time.Date(1904, time.January, 1, 0, 0, 0, 0, time.UTC)
	excelBuggyPeriodStart = time.Date(1900, time.March, 1, 0, 0, 0, -1, time.UTC)
)

// ToTime1900 calculates the floating point number that represents a date for 1900 epoch that counts days.
//
// Excel dates after 28th February 1900 are actually one day out. Excel behaves as though the date
// 29th February 1900 existed, which it didn't. See https://myonlinetraininghub.com/excel-date-and-time
//
// The 1904 Date System is used for compatibility with Excel 2008 for Mac and earlier.
func ToTime1900(t time.Time, date1904 bool) (result float64) {
	if t.Location() != time.UTC {
		t = t.UTC()
	}

	if date1904 {
		return float64(t.Sub(time1904Epoch)) / float64(secondsInADay)
	}

	result = float64(t.Sub(time1900Epoch)) / float64(secondsInADay)
	if t.After(excelBuggyPeriodStart) {
		result += 1.0
	}
	return result
}

// FromTime1900 calculates a time object that represents a floating point number counting days for 1900 epoch.
//
// Excel dates after 28th February 1900 are actually one day out. Excel behaves as though the date
// 29th February 1900 existed, which it didn't. See https://myonlinetraininghub.com/excel-date-and-time
//
// The 1904 Date System is used for compatibility with Excel 2008 for Mac and earlier.
func FromTime1900(t float64, date1904 bool) time.Time {
	if date1904 {
		return time1904Epoch.Add(time.Duration(t * float64(secondsInADay)))
	}

	result := time1900Epoch.Add(time.Duration(t * float64(secondsInADay)))
	if result.After(excelBuggyPeriodStart) {
		result = result.Add(-secondsInADay)
	}
	return result
}
