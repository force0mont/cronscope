package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"cronscope/internal/cron"
)

// ValidationView renders a detailed breakdown of a cron expression's fields.
type ValidationView struct {
	result cron.ValidationResult
	expr   string
}

var (
	fieldLabelStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	fieldValueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	validBadge      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10")).Render("✔ VALID")
	invalidBadge    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("9")).Render("✘ INVALID")
	fieldOrder      = []string{"minute", "hour", "day-of-month", "month", "day-of-week", "descriptor"}
)

// NewValidationView creates a ValidationView for the given expression.
func NewValidationView(expr string) ValidationView {
	return ValidationView{
		expr:   expr,
		result: cron.Validate(expr),
	}
}

// Result returns the underlying ValidationResult.
func (v ValidationView) Result() cron.ValidationResult {
	return v.result
}

// Render produces a human-readable string summarising the validation outcome.
func (v ValidationView) Render() string {
	var sb strings.Builder

	if v.result.Valid {
		sb.WriteString(validBadge)
	} else {
		sb.WriteString(invalidBadge)
	}
	sb.WriteString("  ")
	sb.WriteString(fieldValueStyle.Render(v.expr))
	sb.WriteString("\n")

	if !v.result.Valid {
		sb.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render("  "+v.result.Message))
		sb.WriteString("\n")
		return sb.String()
	}

	if len(v.result.Fields) > 0 {
		sb.WriteString("\n")
		for _, key := range fieldOrder {
			val, ok := v.result.Fields[key]
			if !ok {
				continue
			}
			line := fmt.Sprintf("  %-14s %s\n",
				fieldLabelStyle.Render(key+":"),
				fieldValueStyle.Render(val),
			)
			sb.WriteString(line)
		}
	}
	return sb.String()
}
