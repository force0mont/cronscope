package ui

import (
	"strings"
	"testing"
)

func TestNewScheduleView_ValidExpression(t *testing.T) {
	sv := NewScheduleView("0 * * * *", 5)

	if !sv.IsValid() {
		t.Fatalf("expected valid schedule, got error: %s", sv.Error)
	}
	if len(sv.NextRuns) != 5 {
		t.Errorf("expected 5 next runs, got %d", len(sv.NextRuns))
	}
}

func TestNewScheduleView_InvalidExpression(t *testing.T) {
	sv := NewScheduleView("not-a-cron", 5)

	if sv.IsValid() {
		t.Fatal("expected invalid schedule")
	}
	if sv.Error == "" {
		t.Error("expected non-empty error message")
	}
	if len(sv.NextRuns) != 0 {
		t.Errorf("expected no next runs for invalid expression, got %d", len(sv.NextRuns))
	}
}

func TestScheduleView_Render_Valid(t *testing.T) {
	sv := NewScheduleView("@hourly", 3)
	out := sv.Render()

	if !strings.Contains(out, "@hourly") {
		t.Error("render output should contain the expression")
	}
	if !strings.Contains(out, "Next 3 runs") {
		t.Error("render output should contain next runs header")
	}
	// Each run line starts with a numbered prefix
	if !strings.Contains(out, "1.") {
		t.Error("render output should contain numbered run entries")
	}
}

func TestScheduleView_Render_Invalid(t *testing.T) {
	sv := NewScheduleView("bad expr", 3)
	out := sv.Render()

	if !strings.Contains(out, "Error") {
		t.Error("render output should contain 'Error' label for invalid expression")
	}
	if strings.Contains(out, "Next") {
		t.Error("render output should not contain 'Next' runs section for invalid expression")
	}
}
