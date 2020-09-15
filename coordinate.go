package calculus

import (
	"math"
	"regexp"
	"strconv"
)

var coordinateExp = regexp.MustCompile(`^([A-Z]+)(\d+)$`)

// ParseCoordinate tries to read the row and column information from a string consisting on N english alphabet leters
// followed by a integer number. The letters on the first part represents the columns and the number on the second and
// last part the row (starting on 1).
func ParseCoordinate(coordinate string) (int, int, error) {
	if !coordinateExp.MatchString(coordinate) {
		return -1, -1, NewErrorf(nil, `Invalid coordinate "%s"`, coordinate)
	}
	matches := coordinateExp.FindStringSubmatch(coordinate)

	column := 0
	for i, l := 0, len(matches[1]); i < l; i++ {
		if l > 1 && i < l-1 {
			column += int(math.Pow(26, float64(l-i-1))) * (int(matches[1][i]) - 'A' + 1)
		} else {
			column += int(matches[1][i]) - 'A'
		}
	}
	row, _ := strconv.Atoi(matches[2])

	return row - 1, column, nil
}

// Coordinate calculates the coordinate for a given row and column pair consisting of aN english alphabet letters
// representing the column followed by the row number (starting on 1). The row and column must be non negative integers.
func Coordinate(row, column int) (string, error) {
	if row < 0 || column < 0 {
		return "", NewErrorf(nil, "Row (%d) and column (%d) values must be non negative", row, column)
	}
	letter := ""
	for {
		if column < 26 {
			letter = string('A'+rune(column)) + letter
			break
		} else {
			rest := column % 26
			column = (column / 26) - 1
			letter = string('A'+rune(rest)) + letter
		}
	}
	return letter + strconv.Itoa(row+1), nil
}
