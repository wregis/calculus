package duration

import (
	"regexp"
	"strconv"
	"time"

	"github.com/wregis/calculus"
)

// Adapted from github.com/senseyeio/duration

var durationExp = regexp.MustCompile(`^P((?P<year>\d+)Y)?((?P<month>\d+)M)?((?P<week>\d+)W)?((?P<day>\d+)D)?(T((?P<hour>\d+)H)?((?P<minute>\d+)M)?((?P<second>\d+)S)?)?$`)

// Parse creates a time.Duration fropm an ISO-8601 duration string
func Parse(source string) (time.Duration, error) {
	if !durationExp.MatchString(source) {
		return 0, calculus.NewError(nil, "Could not parse duration string")
	}
	var duration time.Duration
	matches := durationExp.FindStringSubmatch(source)
	for i, name := range durationExp.SubexpNames() {
		part := matches[i]
		if i == 0 || name == "" || part == "" {
			continue
		}
		val, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return duration, err
		}
		switch name {
		case "year":
			duration += time.Duration(val) * time.Hour * 24 * 365 // ok?
		case "month":
			duration += time.Duration(val) * time.Hour * 24 * 30 // ok?
		case "week":
			duration += time.Duration(val) * time.Hour * 24 * 7
		case "day":
			duration += time.Duration(val) * time.Hour * 24
		case "hour":
			duration += time.Duration(val) * time.Hour
		case "minute":
			duration += time.Duration(val) * time.Minute
		case "second":
			duration += time.Duration(val) * time.Second
		default:
			return duration, calculus.NewErrorf(nil, "Unknown field %s", name)
		}
	}

	return duration, nil
}
