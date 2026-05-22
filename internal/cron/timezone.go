package cron

import (
	"fmt"
	"time"
)

// TimezoneInfo holds metadata about a resolved timezone.
type TimezoneInfo struct {
	Name   string
	Offset string
	Loc    *time.Location
}

// ResolveTimezone parses and validates a timezone string, returning a
// TimezoneInfo with the resolved location and its current UTC offset.
// If tzName is empty, the local timezone is used.
func ResolveTimezone(tzName string) (TimezoneInfo, error) {
	if tzName == "" {
		loc := time.Local
		return TimezoneInfo{
			Name:   loc.String(),
			Offset: utcOffset(loc),
			Loc:    loc,
		}, nil
	}

	loc, err := time.LoadLocation(tzName)
	if err != nil {
		return TimezoneInfo{}, fmt.Errorf("unknown timezone %q: %w", tzName, err)
	}

	return TimezoneInfo{
		Name:   loc.String(),
		Offset: utcOffset(loc),
		Loc:    loc,
	}, nil
}

// NextNInTimezone returns the next n scheduled times for expr evaluated in the
// given timezone. If tzName is empty the local timezone is used.
func NextNInTimezone(expr string, n int, tzName string) ([]time.Time, error) {
	tz, err := ResolveTimezone(tzName)
	if err != nil {
		return nil, err
	}

	schedule, err := Parse(expr)
	if err != nil {
		return nil, err
	}

	now := time.Now().In(tz.Loc)
	times := make([]time.Time, 0, n)
	t := now
	for i := 0; i < n; i++ {
		t = schedule.Next(t)
		times = append(times, t)
	}
	return times, nil
}

// utcOffset returns a human-readable UTC offset string for the given location
// at the current moment, e.g. "UTC+05:30" or "UTC-07:00".
func utcOffset(loc *time.Location) string {
	_, offset := time.Now().In(loc).Zone()
	sign := "+"
	if offset < 0 {
		sign = "-"
		offset = -offset
	}
	hours := offset / 3600
	minutes := (offset % 3600) / 60
	return fmt.Sprintf("UTC%s%02d:%02d", sign, hours, minutes)
}
