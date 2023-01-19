// Package timeconv includes different date and time conversions.
package timeconv

import "time"

const (
	hourFormat = "2006-01-02T15:"
	zoneFormat = "Z07:00"
)

// GetNextHalfHour returns next half an hour time.
// For example:
// Arg : 2006-01-02T15:20:40+07:00 > Return: 2006-01-02T15:30:00+07:00
// Arg : 2006-01-02T15:40:40+07:00 > Return: 2006-01-02T16:00:00+07:00
func GetNextHalfHour(t time.Time) time.Time {

	hourStr := t.Format(hourFormat)
	zoneStr := t.Format(zoneFormat)

	timeStr := ""
	sec := t.Minute()*60 + t.Second()

	if sec > 0 && sec <= 1800 {
		timeStr = hourStr + "30:00" + zoneStr
		t, _ = time.Parse(time.RFC3339, timeStr)

	} else if sec > 1800 {
		timeStr = hourStr + "30:00" + zoneStr
		t, _ = time.Parse(time.RFC3339, timeStr)

		// Add 30 min to shift the time.
		t = t.Add(30 * time.Minute)
	} else {
		timeStr = hourStr + "00:00" + zoneStr
		t, _ = time.Parse(time.RFC3339, timeStr)
	}
	return t
}

// GetNextHour returns next  hour time.
// For example:
// Arg : 2006-01-02T15:20:40+07:00 > Return: 2006-01-02T16:00:00+07:00
func GetNextHour(t time.Time) time.Time {

	hourStr := t.Format(hourFormat)
	zoneStr := t.Format(zoneFormat)

	timeStr := ""
	sec := t.Minute()*60 + t.Second()

	if sec < 3600 {
		timeStr = hourStr + "00:00" + zoneStr
		t, _ = time.Parse(time.RFC3339, timeStr)
		t = t.Add(1 * time.Hour)
	}
	return t
}

// SubHalfHour Subtract half an hour.
func SubHalfHour(t time.Time) time.Time {
	return t.Add(-30 * time.Minute)
}

// ToString converts to default time string.
func ToString(t time.Time) string {
	return t.Format(time.RFC3339)
}
