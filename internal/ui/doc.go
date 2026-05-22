// Package ui provides terminal user-interface components for cronscope.
//
// Components
//
// AppModel is the root Bubble Tea model that wires together all sub-components.
//
// InputModel renders the cron-expression text field and surfaces validation
// errors inline.
//
// ValidationView renders a colour-coded summary of each cron field (minute,
// hour, day-of-month, month, day-of-week).
//
// ScheduleView renders the next N upcoming run times for a validated
// expression.
//
// HistoryModel renders a navigable list of recently used expressions and
// emits a HistorySelectedMsg when the user selects one.  Persistence is
// handled by internal/cron.History; the UI layer only consumes the
// in-memory slice of expression strings.
//
// Key bindings (AppModel)
//
//   h          toggle history panel
//   ctrl+c / q quit
package ui
