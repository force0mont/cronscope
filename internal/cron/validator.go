package cron

import (
	"fmt"
	"strconv"
	"strings"
)

// ValidationResult holds the outcome of validating a cron expression.
type ValidationResult struct {
	Valid   bool
	Message string
	Fields  map[string]string
}

// fieldSpec describes allowed range for a cron field.
type fieldSpec struct {
	name string
	min  int
	max  int
}

var standardFields = []fieldSpec{
	{"minute", 0, 59},
	{"hour", 0, 23},
	{"day-of-month", 1, 31},
	{"month", 1, 12},
	{"day-of-week", 0, 7},
}

// Validate checks a cron expression for structural correctness and returns
// a ValidationResult with field-level detail.
func Validate(expr string) ValidationResult {
	expr = strings.TrimSpace(expr)
	if strings.HasPrefix(expr, "@") {
		_, err := Parse(expr)
		if err != nil {
			return ValidationResult{Valid: false, Message: err.Error()}
		}
		return ValidationResult{Valid: true, Message: "descriptor expression", Fields: map[string]string{"descriptor": expr}}
	}

	parts := strings.Fields(expr)
	if len(parts) != 5 {
		return ValidationResult{
			Valid:   false,
			Message: fmt.Sprintf("expected 5 fields, got %d", len(parts)),
		}
	}

	fields := make(map[string]string, 5)
	for i, spec := range standardFields {
		if err := validateField(parts[i], spec); err != nil {
			return ValidationResult{
				Valid:   false,
				Message: fmt.Sprintf("%s: %s", spec.name, err.Error()),
				Fields:  fields,
			}
		}
		fields[spec.name] = parts[i]
	}
	return ValidationResult{Valid: true, Message: "valid expression", Fields: fields}
}

func validateField(field string, spec fieldSpec) error {
	if field == "*" {
		return nil
	}
	// Handle step values e.g. */5 or 1-5/2
	parts := strings.SplitN(field, "/", 2)
	if len(parts) == 2 {
		step, err := strconv.Atoi(parts[1])
		if err != nil || step < 1 {
			return fmt.Errorf("invalid step value %q", parts[1])
		}
		if parts[0] == "*" {
			return nil
		}
		field = parts[0]
	}
	// Handle ranges e.g. 1-5
	if strings.Contains(field, "-") {
		bounds := strings.SplitN(field, "-", 2)
		lo, err1 := strconv.Atoi(bounds[0])
		hi, err2 := strconv.Atoi(bounds[1])
		if err1 != nil || err2 != nil {
			return fmt.Errorf("invalid range %q", field)
		}
		if lo < spec.min || hi > spec.max || lo > hi {
			return fmt.Errorf("range %d-%d out of bounds [%d-%d]", lo, hi, spec.min, spec.max)
		}
		return nil
	}
	// Handle lists e.g. 1,3,5
	for _, part := range strings.Split(field, ",") {
		v, err := strconv.Atoi(part)
		if err != nil {
			return fmt.Errorf("invalid value %q", part)
		}
		if v < spec.min || v > spec.max {
			return fmt.Errorf("value %d out of bounds [%d-%d]", v, spec.min, spec.max)
		}
	}
	return nil
}
