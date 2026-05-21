package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/user/cronscope/internal/cron"
)

// ScheduleView holds the state for rendering a cron schedule preview.
type ScheduleView struct {
	Expression string
	Count      int
	NextRuns   []time.Time
	Error      string
}

// NewScheduleView creates a ScheduleView by parsing the given expression
// and computing the next n upcoming run times.
func NewScheduleView(expr string, count int) *ScheduleView {
	sv := &ScheduleView{
		Expression: expr,
		Count:      count,
	}

	sch, err := cron.Parse(expr)
	if err != nil {
		sv.Error = fmt.Sprintf("invalid expression: %v", err)
		return sv
	}

	sv.NextRuns = cron.NextN(sch, time.Now(), count)
	return sv
}

// IsValid returns true when no parse error occurred.
func (sv *ScheduleView) IsValid() bool {
	return sv.Error == ""
}

// Render returns a plain-text representation of the schedule preview.
func (sv *ScheduleView) Render() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Expression : %s\n", sv.Expression))

	if !sv.IsValid() {
		sb.WriteString(fmt.Sprintf("Error      : %s\n", sv.Error))
		return sb.String()
	}

	sb.WriteString(fmt.Sprintf("Next %d runs:\n", sv.Count))
	for i, t := range sv.NextRuns {
		sb.WriteString(fmt.Sprintf("  %2d. %s\n", i+1, t.Format("2006-01-02 15:04:05 MST")))
	}

	return sb.String()
}
