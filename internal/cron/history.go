package cron

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// HistoryEntry represents a previously used cron expression.
type HistoryEntry struct {
	Expression string    `json:"expression"`
	UsedAt     time.Time `json:"used_at"`
}

// History manages a persisted list of recently used cron expressions.
type History struct {
	Entries []HistoryEntry `json:"entries"`
	MaxSize int            `json:"-"`
	path    string
}

// NewHistory creates a History backed by the given file path.
func NewHistory(path string, maxSize int) *History {
	return &History{path: path, MaxSize: maxSize}
}

// Load reads history from disk. Missing file is treated as empty history.
func (h *History) Load() error {
	data, err := os.ReadFile(h.path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(data, h)
}

// Save persists history to disk, creating parent directories as needed.
func (h *History) Save() error {
	if err := os.MkdirAll(filepath.Dir(h.path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(h.path, data, 0o644)
}

// Add inserts an expression at the front, deduplicates, and trims to MaxSize.
func (h *History) Add(expr string) {
	filtered := make([]HistoryEntry, 0, len(h.Entries))
	for _, e := range h.Entries {
		if e.Expression != expr {
			filtered = append(filtered, e)
		}
	}
	entry := HistoryEntry{Expression: expr, UsedAt: time.Now().UTC()}
	h.Entries = append([]HistoryEntry{entry}, filtered...)
	if h.MaxSize > 0 && len(h.Entries) > h.MaxSize {
		h.Entries = h.Entries[:h.MaxSize]
	}
}

// Expressions returns the stored expressions in recency order.
func (h *History) Expressions() []string {
	out := make([]string, len(h.Entries))
	for i, e := range h.Entries {
		out[i] = e.Expression
	}
	return out
}

// Clear removes all entries from the history.
func (h *History) Clear() error {
	h.Entries = nil
	return h.Save()
}
