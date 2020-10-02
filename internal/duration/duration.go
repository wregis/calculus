package duration

import (
	"bytes"
	"regexp"
	"strconv"
	"text/template"
	"time"

	"github.com/wregis/calculus/internal/errors"
)

// Adapted from github.com/senseyeio/duration

var durationExp = regexp.MustCompile(`^P((?P<year>\d+)Y)?((?P<month>\d+)M)?((?P<week>\d+)W)?((?P<day>\d+)D)?(T((?P<hour>\d+)H)?((?P<minute>\d+)M)?((?P<second>\d+)S)?)?$`)

// Parse creates a time.Duration fropm an ISO-8601 duration string.
func Parse(source string) (time.Duration, error) {
	if !durationExp.MatchString(source) {
		return 0, errors.New(nil, "Could not parse duration string")
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
			return duration, errors.Newf(nil, "Unknown field %s", name)
		}
	}

	return duration, nil
}

var durationTmpl = template.Must(template.New("duration").Parse(
	`P{{if .Y}}{{.Y}}Y{{end}}{{if .M}}{{.M}}M{{end}}{{if .W}}{{.W}}W{{end}}{{if .D}}{{.D}}D{{end}}{{if .HasTimePart}}` +
		`T{{end }}{{if .TH}}{{.TH}}H{{end}}{{if .TM}}{{.TM}}M{{end}}{{if .TS}}{{.TS}}S{{end}}`,
))

// Format returns an ISO-8601-ish representation of the duration.
func Format(duration time.Duration) string {
	if duration == 0 {
		return "P0D"
	}
	var d struct {
		Y, M, W, D, TH, TM, TS uint
		HasTimePart            bool
	}
	d.TM = uint(duration.Seconds()) / 60
	d.TS = uint(duration.Seconds()) % 60
	if d.TM >= 60 {
		d.TH = d.TM / 60
		d.TM = d.TM % 60
	}
	if d.TH >= 24 {
		d.D = d.TH / 24
		d.TH = d.TH % 24
	}
	if d.D >= 365 {
		d.Y = d.D / 365
		d.D = d.D % 365
	}
	if d.D >= 30 {
		d.M = d.D / 30
		d.D = d.D % 30
	}
	if d.D >= 7 {
		d.W = d.D / 7
		d.D = d.D % 7
	}

	if d.TS > 0 || d.TM > 0 || d.TH > 0 {
		d.HasTimePart = true
	}
	var s bytes.Buffer
	if err := durationTmpl.Execute(&s, d); err != nil {
		return ""
	}
	return s.String()
}
