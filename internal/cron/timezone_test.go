package cron

import (
	"testing"
	"time"
)

func TestResolveTimezone_EmptyUsesLocal(t *testing.T) {
	info, err := ResolveTimezone("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if info.Loc == nil {
		t.Fatal("expected non-nil location")
	}
	if info.Name == "" {
		t.Fatal("expected non-empty timezone name")
	}
}

func TestResolveTimezone_ValidZone(t *testing.T) {
	info, err := ResolveTimezone("UTC")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if info.Name != "UTC" {
		t.Errorf("expected UTC, got %q", info.Name)
	}
	if info.Offset != "UTC+00:00" {
		t.Errorf("expected UTC+00:00, got %q", info.Offset)
	}
}

func TestResolveTimezone_InvalidZone(t *testing.T) {
	_, err := ResolveTimezone("Not/AReal/Zone")
	if err == nil {
		t.Fatal("expected error for invalid timezone")
	}
}

func TestNextNInTimezone_ReturnsCorrectCount(t *testing.T) {
	times, err := NextNInTimezone("* * * * *", 5, "UTC")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(times) != 5 {
		t.Errorf("expected 5 times, got %d", len(times))
	}
}

func TestNextNInTimezone_TimesAreInRequestedZone(t *testing.T) {
	const tz = "America/New_York"
	loc, err := time.LoadLocation(tz)
	if err != nil {
		t.Skip("timezone data not available")
	}
	times, err := NextNInTimezone("0 9 * * *", 3, tz)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, ts := range times {
		if ts.Location().String() != loc.String() {
			t.Errorf("expected location %s, got %s", loc, ts.Location())
		}
	}
}

func TestNextNInTimezone_InvalidExpr(t *testing.T) {
	_, err := NextNInTimezone("not-a-cron", 3, "UTC")
	if err == nil {
		t.Fatal("expected error for invalid expression")
	}
}

func TestNextNInTimezone_InvalidTimezone(t *testing.T) {
	_, err := NextNInTimezone("* * * * *", 3, "Fake/Zone")
	if err == nil {
		t.Fatal("expected error for invalid timezone")
	}
}
