package cron

import (
	"testing"
)

func TestValidate_ValidStandardExpression(t *testing.T) {
	res := Validate("0 9 * * 1")
	if !res.Valid {
		t.Fatalf("expected valid, got: %s", res.Message)
	}
	if res.Fields["minute"] != "0" {
		t.Errorf("expected minute=0, got %q", res.Fields["minute"])
	}
	if res.Fields["day-of-week"] != "1" {
		t.Errorf("expected day-of-week=1, got %q", res.Fields["day-of-week"])
	}
}

func TestValidate_TooFewFields(t *testing.T) {
	res := Validate("0 9 * *")
	if res.Valid {
		t.Fatal("expected invalid for 4-field expression")
	}
}

func TestValidate_OutOfRangeMinute(t *testing.T) {
	res := Validate("60 9 * * *")
	if res.Valid {
		t.Fatal("expected invalid for minute=60")
	}
	if res.Message == "" {
		t.Error("expected non-empty error message")
	}
}

func TestValidate_OutOfRangeHour(t *testing.T) {
	res := Validate("0 24 * * *")
	if res.Valid {
		t.Fatalf("expected invalid for hour=24, got: %s", res.Message)
	}
}

func TestValidate_StepExpression(t *testing.T) {
	res := Validate("*/15 * * * *")
	if !res.Valid {
		t.Fatalf("expected valid step expression, got: %s", res.Message)
	}
}

func TestValidate_RangeExpression(t *testing.T) {
	res := Validate("0 9-17 * * 1-5")
	if !res.Valid {
		t.Fatalf("expected valid range expression, got: %s", res.Message)
	}
}

func TestValidate_ListExpression(t *testing.T) {
	res := Validate("0 8,12,18 * * *")
	if !res.Valid {
		t.Fatalf("expected valid list expression, got: %s", res.Message)
	}
}

func TestValidate_Descriptor(t *testing.T) {
	res := Validate("@daily")
	if !res.Valid {
		t.Fatalf("expected valid descriptor, got: %s", res.Message)
	}
	if res.Fields["descriptor"] != "@daily" {
		t.Errorf("expected descriptor field, got %v", res.Fields)
	}
}

func TestValidate_InvalidDescriptor(t *testing.T) {
	res := Validate("@nope")
	if res.Valid {
		t.Fatal("expected invalid for unknown descriptor")
	}
}

func TestValidate_InvalidStep(t *testing.T) {
	res := Validate("*/0 * * * *")
	if res.Valid {
		t.Fatal("expected invalid for step=0")
	}
}
