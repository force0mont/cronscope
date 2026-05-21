package ui

import (
	"errors"
	"strings"
	"testing"
)

func TestNewInputModel_DefaultState(t *testing.T) {
	m := NewInputModel()

	if m.Value() != "" {
		t.Errorf("expected empty value, got %q", m.Value())
	}
	if m.Err != nil {
		t.Errorf("expected nil error, got %v", m.Err)
	}
}

func TestInputModel_SetError_DisplayedInView(t *testing.T) {
	m := NewInputModel()
	sentinel := errors.New("invalid cron expression")
	m.SetError(sentinel)

	view := m.View()
	if !strings.Contains(view, sentinel.Error()) {
		t.Errorf("expected view to contain error %q, got:\n%s", sentinel.Error(), view)
	}
}

func TestInputModel_Reset_ClearsError(t *testing.T) {
	m := NewInputModel()
	m.SetError(errors.New("some error"))
	m.Reset()

	if m.Err != nil {
		t.Errorf("expected nil error after Reset, got %v", m.Err)
	}
	if m.Value() != "" {
		t.Errorf("expected empty value after Reset, got %q", m.Value())
	}
}

func TestInputModel_View_ContainsPlaceholderHint(t *testing.T) {
	m := NewInputModel()
	view := m.View()

	if !strings.Contains(view, "cron expression") {
		t.Errorf("expected view to contain 'cron expression', got:\n%s", view)
	}
}

func TestInputModel_NoError_ViewOmitsWarning(t *testing.T) {
	m := NewInputModel()
	view := m.View()

	if strings.Contains(view, "⚠") {
		t.Errorf("expected no warning symbol when no error, got:\n%s", view)
	}
}
