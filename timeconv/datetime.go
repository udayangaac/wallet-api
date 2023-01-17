// Package timeconv includes different date and time conversions.
package timefmt

import "time"

const (
	hourFormat = "2006-01-02T15:"
	zoneFormat = "Z07:00"
)

// GetPreviousHalfHour returns previous latest half an hour time.
// For example:
// Arg : 2006-01-02T15:20:40+07:00 > Return: 2006-01-02T15:00:00+07:00
// Arg : 2006-01-02T15:40:40+07:00 > Return: 2006-01-02T15:30:00+07:00
func GetPreviousHalfHour(t time.Time) time.Time {

	hourStr := t.Format(hourFormat)
	zoneStr := t.Format(zoneFormat)

	timeStr := ""
	if t.Minute() >= 30 {
		timeStr = hourStr + "30:00" + zoneStr
	} else {
		timeStr = hourStr + "00:00" + zoneStr
	}

	// Error can be neglected because formats of both arguments do not change.
	t, _ = time.Parse(time.RFC3339, timeStr)
	return t
}
