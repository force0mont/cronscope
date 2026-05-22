package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"cronscope/internal/cron"
)

var (
	intervalHeaderStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("33"))
	intervalRowStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	intervalGapStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
)

// IntervalView renders the gaps between consecutive cron fire times.
type IntervalView struct {
	expr      string
	ref       time.Time
	count     int
	err       error
}

// NewIntervalView creates an IntervalView for the given expression.
func NewIntervalView(expr string, ref time.Time, count int) IntervalView {
	return IntervalView{
		expr:  expr,
		ref:   ref,
		count: count,
	}
}

// Render returns the string representation of the interval view.
func (v IntervalView) Render() string {
	if v.expr == "" {
		return ""
	}

	n := v.count
	if n < 2 {
		n = 6
	}

	times, err := cron.NextN(v.expr, v.ref, n)
	if err != nil {
		return fmt.Sprintf("interval error: %v", err)
	}

	intervals, err := cron.NextIntervals(v.expr, v.ref, n)
	if err != nil {
		return fmt.Sprintf("interval error: %v", err)
	}

	var sb strings.Builder
	sb.WriteString(intervalHeaderStyle.Render("Run Intervals") + "\n")

	for i, iv := range intervals {
		runLabel := intervalRowStyle.Render(fmt.Sprintf("  %s → %s",
			times[i].Format("15:04:05"),
			times[i+1].Format("15:04:05"),
		))
		gapLabel := intervalGapStyle.Render(fmt.Sprintf(" (+%s)", iv.Label))
		sb.WriteString(runLabel + gapLabel + "\n")
	}

	return sb.String()
}
