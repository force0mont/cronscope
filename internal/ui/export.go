// Package ui provides terminal UI components for cronscope.
package ui

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/user/cronscope/internal/cron"
)

// ExportView renders a formatted export of scheduled run times
// for a given cron expression, suitable for copying or saving.
type ExportView struct {
	expr  string
	count int
	times []time.Time
	err   error
}

// NewExportView creates an ExportView for the given cron expression,
// pre-computing the next n run times.
func NewExportView(expr string, n int) *ExportView {
	v := &ExportView{expr: expr, count: n}
	times, err := cron.NextN(expr, n)
	if err != nil {
		v.err = err
	} else {
		v.times = times
	}
	return v
}

// Render returns a human-readable string representation of the export,
// showing the expression and its upcoming run times.
func (v *ExportView) Render() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Expression: %s\n", v.expr))

	if v.err != nil {
		sb.WriteString(fmt.Sprintf("Error: %s\n", v.err.Error()))
		return sb.String()
	}

	sb.WriteString(fmt.Sprintf("Next %d runs:\n", v.count))
	for i, t := range v.times {
		sb.WriteString(fmt.Sprintf("  %d. %s\n", i+1, t.Format(time.RFC3339)))
	}

	return sb.String()
}

// exportPayload is the JSON structure for exported cron data.
type exportPayload struct {
	Expression string   `json:"expression"`
	NextRuns   []string `json:"next_runs"`
	Error      string   `json:"error,omitempty"`
}

// FormatJSON returns a JSON-encoded string of the export data,
// including the expression and next run timestamps.
func (v *ExportView) FormatJSON() string {
	payload := exportPayload{
		Expression: v.expr,
	}

	if v.err != nil {
		payload.Error = v.err.Error()
	} else {
		runs := make([]string, len(v.times))
		for i, t := range v.times {
			runs[i] = t.Format(time.RFC3339)
		}
		payload.NextRuns = runs
	}

	b, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	}
	return string(b)
}
