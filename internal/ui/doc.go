// Package ui provides terminal UI components for cronscope.
//
// It builds on top of the cron parsing primitives in internal/cron to offer
// higher-level views that can be embedded in an interactive TUI or used for
// plain-text output.
//
// # ScheduleView
//
// ScheduleView is the primary component. Given a cron expression and a desired
// preview count it parses the expression, computes the next N run times, and
// exposes a Render method that returns a human-readable summary:
//
//	// Print the next 10 runs for a cron expression.
//	sv := ui.NewScheduleView("0 9 * * 1-5", 10)
//	if !sv.IsValid() {
//		log.Fatal(sv.Error)
//	}
//	fmt.Print(sv.Render())
//
// Future components (input fields, key-bindings, colour themes) will be added
// to this package as the project matures.
package ui
