package cron

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// ParseResult holds the result of parsing a cron expression.
type ParseResult struct {
	Expression string
	Schedule   cron.Schedule
	Valid      bool
	Error      string
}

// NextRuns returns the next n scheduled times after the given start time.
type NextRuns struct {
	Expression string
	Runs       []time.Time
}

// Parse validates and parses a cron expression.
// Supports standard 5-field expressions and optional seconds field.
func Parse(expr string) ParseResult {
	parser := cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	schedule, err := parser.Parse(expr)
	if err != nil {
		return ParseResult{
			Expression: expr,
			Valid:      false,
			Error:      fmt.Sprintf("invalid expression: %v", err),
		}
	}

	return ParseResult{
		Expression: expr,
		Schedule:   schedule,
		Valid:      true,
	}
}

// NextN computes the next n run times for a valid cron expression
// starting from the given time.
func NextN(expr string, from time.Time, n int) (NextRuns, error) {
	result := Parse(expr)
	if !result.Valid {
		return NextRuns{}, fmt.Errorf(result.Error)
	}

	runs := make([]time.Time, 0, n)
	t := from
	for i := 0; i < n; i++ {
		t = result.Schedule.Next(t)
		runs = append(runs, t)
	}

	return NextRuns{
		Expression: expr,
		Runs:       runs,
	}, nil
}
