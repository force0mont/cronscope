package cron_test

import (
	"testing"
	"time"

	"github.com/indaco/cronscope/internal/cron"
)

func TestParse_ValidExpression(t *testing.T) {
	result := cron.Parse("0 9 * * 1-5")
	if !result.Valid {
		t.Fatalf("expected valid expression, got error: %s", result.Error)
	}
	if result.Schedule == nil {
		t.Fatal("expected non-nil schedule")
	}
}

func TestParse_InvalidExpression(t *testing.T) {
	result := cron.Parse("not a cron")
	if result.Valid {
		t.Fatal("expected invalid expression")
	}
	if result.Error == "" {
		t.Fatal("expected non-empty error message")
	}
}

func TestParse_Descriptor(t *testing.T) {
	result := cron.Parse("@daily")
	if !result.Valid {
		t.Fatalf("expected @daily to be valid, got: %s", result.Error)
	}
}

func TestNextN_ReturnsCorrectCount(t *testing.T) {
	from := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	runs, err := cron.NextN("0 9 * * *", from, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(runs.Runs) != 5 {
		t.Fatalf("expected 5 runs, got %d", len(runs.Runs))
	}
}

func TestNextN_RunsAreAscending(t *testing.T) {
	from := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	runs, err := cron.NextN("*/15 * * * *", from, 4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for i := 1; i < len(runs.Runs); i++ {
		if !runs.Runs[i].After(runs.Runs[i-1]) {
			t.Errorf("run %d (%v) is not after run %d (%v)", i, runs.Runs[i], i-1, runs.Runs[i-1])
		}
	}
}

func TestNextN_InvalidExpression(t *testing.T) {
	from := time.Now()
	_, err := cron.NextN("bad expr", from, 3)
	if err == nil {
		t.Fatal("expected error for invalid expression")
	}
}
