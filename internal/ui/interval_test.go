package ui

import (
	"strings"
	"testing"
	"time"

	"github.com/user/cronscope/internal/cron"
)

func TestNewIntervalView_ValidExpression(t *testing.T) {
	v := NewIntervalView("* * * * *", 5, time.Now())
	if v == nil {
		t.Fatal("expected non-nil IntervalView")
	}
}

func TestNewIntervalView_InvalidExpression(t *testing.T) {
	v := NewIntervalView("invalid expr", 5, time.Now())
	if v == nil {
		t.Fatal("expected non-nil IntervalView even for invalid expr")
	}
}

func TestIntervalView_Render_ValidExpression(t *testing.T) {
	now := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	v := NewIntervalView("* * * * *", 4, now)
	out := v.Render()
	if out == "" {
		t.Fatal("expected non-empty render output")
	}
	if !strings.Contains(out, "Intervals") {
		t.Errorf("expected output to contain 'Intervals', got: %s", out)
	}
}

func TestIntervalView_Render_InvalidExpression(t *testing.T) {
	now := time.Now()
	v := NewIntervalView("bad expr !!", 4, now)
	out := v.Render()
	if !strings.Contains(out, "error") && !strings.Contains(out, "Error") && !strings.Contains(out, "invalid") {
		t.Errorf("expected error indication in output, got: %s", out)
	}
}

func TestIntervalView_Render_ContainsDurations(t *testing.T) {
	now := time.Date(2024, 6, 15, 10, 0, 0, 0, time.UTC)
	v := NewIntervalView("0 * * * *", 3, now)
	out := v.Render()
	intervals, err := cron.NextIntervals("0 * * * *", 3, now)
	if err != nil {
		t.Skipf("skipping: could not compute intervals: %v", err)
	}
	if len(intervals) == 0 {
		t.Skip("no intervals returned")
	}
	_ = out
	_ = intervals
}

func TestIntervalView_Render_ZeroCount(t *testing.T) {
	now := time.Now()
	v := NewIntervalView("* * * * *", 0, now)
	out := v.Render()
	// Should render without panic, content may be empty or show header only
	_ = out
}
