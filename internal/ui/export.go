package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/user/cronscope/internal/cron"
)

// ExportView renders a plain-text exportable summary of a cron expression,
// including the next N scheduled times and their intervals.
type ExportView struct {
	expr  string
	count int
	from  time.Time
}

// NewExportView creates an ExportView for the given expression.
func NewExportView(expr string, count int, from time.Time) *ExportView {
	return &ExportView{
		expr:  expr,
		count: count,
		from:  from,
	}
}

// Render returns a plain-text string suitable for copying or saving.
func (e *ExportView) Render() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Cron Expression : %s\n", e.expr))
	sb.WriteString(fmt.Sprintf("Generated At    : %s\n", e.from.Format(time.RFC3339)))
	sb.WriteString(strings.Repeat("-", 40) + "\n")

	times, err := cron.NextN(e.expr, e.count, e.from)
	if err != nil {
		sb.WriteString(fmt.Sprintf("Error: %v\n", err))
		return sb.String()
	}

	if len(times) == 0 {
		sb.WriteString("No upcoming runs found.\n")
		return sb.String()
	}

	sb.WriteString("Next Scheduled Runs:\n")
	for i, t := range times {
		sb.WriteString(fmt.Sprintf("  %2d. %s\n", i+1, t.Format("2006-01-02 15:04:05 MST")))
	}

	sb.WriteString(strings.Repeat("-", 40) + "\n")

	intervals, err := cron.NextIntervals(e.expr, e.count, e.from)
	if err == nil && len(intervals) > 0 {
		sb.WriteString("Intervals Between Runs:\n")
		for i, d := range intervals {
			sb.WriteString(fmt.Sprintf("  Run %d → %d : %s\n", i+1, i+2, d.Round(time.Second)))
		}
	}

	return sb.String()
}
