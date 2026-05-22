package ui

import (
	"strings"
	"testing"
)

func TestNewExportView_ValidExpression(t *testing.T) {
	v := NewExportView("0 9 * * 1-5", 5)
	if v == nil {
		t.Fatal("expected non-nil ExportView")
	}
}

func TestNewExportView_InvalidExpression(t *testing.T) {
	v := NewExportView("invalid expr", 5)
	if v == nil {
		t.Fatal("expected non-nil ExportView even for invalid expr")
	}
}

func TestExportView_Render_ValidExpression(t *testing.T) {
	v := NewExportView("0 9 * * 1-5", 3)
	out := v.Render()
	if out == "" {
		t.Fatal("expected non-empty render output")
	}
}

func TestExportView_Render_InvalidExpression(t *testing.T) {
	v := NewExportView("bad expr !!", 3)
	out := v.Render()
	if !strings.Contains(out, "error") && !strings.Contains(out, "Error") && !strings.Contains(out, "invalid") {
		t.Errorf("expected error indication in render output, got: %s", out)
	}
}

func TestExportView_Render_ContainsExpression(t *testing.T) {
	expr := "*/15 * * * *"
	v := NewExportView(expr, 4)
	out := v.Render()
	if !strings.Contains(out, expr) {
		t.Errorf("expected render output to contain expression %q, got: %s", expr, out)
	}
}

func TestExportView_Render_ContainsTimestamps(t *testing.T) {
	v := NewExportView("0 0 * * *", 3)
	out := v.Render()
	// Should contain at least one date-like string (year)
	if !strings.Contains(out, "202") {
		t.Errorf("expected render output to contain timestamps, got: %s", out)
	}
}

func TestExportView_FormatJSON_ContainsFields(t *testing.T) {
	v := NewExportView("0 12 * * *", 2)
	json := v.FormatJSON()
	if !strings.Contains(json, "expression") {
		t.Errorf("expected JSON to contain 'expression' field, got: %s", json)
	}
	if !strings.Contains(json, "next_runs") {
		t.Errorf("expected JSON to contain 'next_runs' field, got: %s", json)
	}
}
