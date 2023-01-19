package timeconv

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNextHalfHour(t *testing.T) {

	t1, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	t1Expected, _ := time.Parse(time.RFC3339, "2006-01-02T15:30:00+07:00")

	t2, _ := time.Parse(time.RFC3339, "2020-01-02T15:34:05+07:00")
	t2Expected, _ := time.Parse(time.RFC3339, "2020-01-02T16:00:00+07:00")

	t3, _ := time.Parse(time.RFC3339, "2020-01-02T15:00:00+07:00")
	t3Expected, _ := time.Parse(time.RFC3339, "2020-01-02T15:00:00+07:00")

	tests := []struct {
		Time     time.Time
		Expected time.Time
	}{
		{
			Time:     t1,
			Expected: t1Expected,
		},
		{
			Time:     t2,
			Expected: t2Expected,
		},
		{
			Time:     t3,
			Expected: t3Expected,
		},
	}

	for _, test := range tests {
		actual := GetNextHalfHour(test.Time)
		assert.Equal(t, test.Expected, actual)
	}
}

func TestGetNextHour(t *testing.T) {

	t1, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	t1Expected, _ := time.Parse(time.RFC3339, "2006-01-02T16:00:00+07:00")

	t2, _ := time.Parse(time.RFC3339, "2020-01-02T15:34:05+07:00")
	t2Expected, _ := time.Parse(time.RFC3339, "2020-01-02T16:00:00+07:00")

	t3, _ := time.Parse(time.RFC3339, "2020-01-02T16:00:00+07:00")
	t3Expected, _ := time.Parse(time.RFC3339, "2020-01-02T16:00:00+07:00")

	tests := []struct {
		Time     time.Time
		Expected time.Time
	}{
		{
			Time:     t1,
			Expected: t1Expected,
		},
		{
			Time:     t2,
			Expected: t2Expected,
		},
		{
			Time:     t3,
			Expected: t3Expected,
		},
	}

	for _, test := range tests {
		actual := GetNextHour(test.Time)
		assert.Equal(t, test.Expected, actual)
	}
}

func TestSubHalfHour(t *testing.T) {

	t1, _ := time.Parse(time.RFC3339, "2006-01-02T15:00:00+07:00")
	t1Expected, _ := time.Parse(time.RFC3339, "2006-01-02T14:30:00+07:00")

	tests := []struct {
		Time     time.Time
		Expected time.Time
	}{
		{
			Time:     t1,
			Expected: t1Expected,
		},
	}

	for _, test := range tests {
		actual := SubHalfHour(test.Time)
		assert.Equal(t, test.Expected, actual)
	}
}

func TestSubToString(t *testing.T) {

	t1Str := "2006-01-02T15:00:05+07:00"
	t1, _ := time.Parse(time.RFC3339, "2006-01-02T15:00:05+07:00")

	tests := []struct {
		Time     time.Time
		Expected string
	}{
		{
			Time:     t1,
			Expected: t1Str,
		},
	}

	for _, test := range tests {
		actual := ToString(test.Time)
		assert.Equal(t, test.Expected, actual)
	}
}
