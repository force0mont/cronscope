package cron

import (
	"fmt"
	"time"
)

// Interval represents the human-readable gap between two cron fire times.
type Interval struct {
	Duration time.Duration
	Label    string
}

// NextIntervals returns the durations between consecutive next-run times
// for the given cron expression, starting from the reference time t.
// It returns n-1 intervals for n next runs.
func NextIntervals(expr string, t time.Time, n int) ([]Interval, error) {
	if n < 2 {
		return nil, fmt.Errorf("n must be at least 2 to compute intervals")
	}

	times, err := NextN(expr, t, n)
	if err != nil {
		return nil, err
	}

	intervals := make([]Interval, len(times)-1)
	for i := 1; i < len(times); i++ {
		d := times[i].Sub(times[i-1])
		intervals[i-1] = Interval{
			Duration: d,
			Label:    formatDuration(d),
		}
	}
	return intervals, nil
}

// formatDuration converts a duration to a short human-readable string.
func formatDuration(d time.Duration) string {
	switch {
	case d < time.Minute:
		return fmt.Sprintf("%ds", int(d.Seconds()))
	case d < time.Hour:
		m := int(d.Minutes())
		s := int(d.Seconds()) % 60
		if s == 0 {
			return fmt.Sprintf("%dm", m)
		}
		return fmt.Sprintf("%dm %ds", m, s)
	case d < 24*time.Hour:
		h := int(d.Hours())
		m := int(d.Minutes()) % 60
		if m == 0 {
			return fmt.Sprintf("%dh", h)
		}
		return fmt.Sprintf("%dh %dm", h, m)
	default:
		days := int(d.Hours()) / 24
		h := int(d.Hours()) % 24
		if h == 0 {
			return fmt.Sprintf("%dd", days)
		}
		return fmt.Sprintf("%dd %dh", days, h)
	}
}
