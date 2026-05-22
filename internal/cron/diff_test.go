package cron

import (
	"testing"
	"time"
)

func TestNextIntervals_ReturnsNMinusOne(t *testing.T) {
	ref := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	intervals, err := NextIntervals("* * * * *", ref, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(intervals) != 4 {
		t.Errorf("expected 4 intervals, got %d", len(intervals))
	}
}

func TestNextIntervals_EachOneMinuteApart(t *testing.T) {
	ref := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	intervals, err := NextIntervals("* * * * *", ref, 4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for i, iv := range intervals {
		if iv.Duration != time.Minute {
			t.Errorf("interval[%d]: expected 1m, got %v", i, iv.Duration)
		}
	}
}

func TestNextIntervals_InvalidN(t *testing.T) {
	ref := time.Now()
	_, err := NextIntervals("* * * * *", ref, 1)
	if err == nil {
		t.Error("expected error for n=1, got nil")
	}
}

func TestNextIntervals_InvalidExpr(t *testing.T) {
	ref := time.Now()
	_, err := NextIntervals("invalid expr", ref, 3)
	if err == nil {
		t.Error("expected error for invalid expression")
	}
}

func TestFormatDuration_Seconds(t *testing.T) {
	result := formatDuration(30 * time.Second)
	if result != "30s" {
		t.Errorf("expected '30s', got %q", result)
	}
}

func TestFormatDuration_Minutes(t *testing.T) {
	result := formatDuration(5 * time.Minute)
	if result != "5m" {
		t.Errorf("expected '5m', got %q", result)
	}
}

func TestFormatDuration_Hours(t *testing.T) {
	result := formatDuration(2 * time.Hour)
	if result != "2h" {
		t.Errorf("expected '2h', got %q", result)
	}
}

func TestFormatDuration_Days(t *testing.T) {
	result := formatDuration(48 * time.Hour)
	if result != "2d" {
		t.Errorf("expected '2d', got %q", result)
	}
}
