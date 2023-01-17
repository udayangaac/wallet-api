package timefmt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNearestHalfHour(t *testing.T) {

	t1, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	t1Expected, _ := time.Parse(time.RFC3339, "2006-01-02T15:00:00+07:00")

	t2, _ := time.Parse(time.RFC3339, "2020-01-02T15:34:05+07:00")
	t2Expected, _ := time.Parse(time.RFC3339, "2020-01-02T15:30:00+07:00")

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
	}

	for _, test := range tests {
		actual := GetPreviousHalfHour(test.Time)
		assert.Equal(t, test.Expected, actual)
	}
}
