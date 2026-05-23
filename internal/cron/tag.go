package cron

import (
	"fmt"
	"sort"
	"strings"
)

// Tag represents a user-defined label attached to a cron expression.
type Tag struct {
	Label      string `json:"label"`
	Expression string `json:"expression"`
}

// TagStore manages a collection of tags keyed by label.
type TagStore struct {
	tags map[string][]string // label -> []expression
	path string
}

// NewTagStore creates an in-memory tag store backed by the given file path.
func NewTagStore(path string) *TagStore {
	return &TagStore{
		tags: make(map[string][]string),
		path: path,
	}
}

// Add associates a label with an expression. Duplicate pairs are ignored.
func (ts *TagStore) Add(label, expression string) error {
	label = strings.TrimSpace(label)
	if label == "" {
		return fmt.Errorf("tag label must not be empty")
	}
	for _, existing := range ts.tags[label] {
		if existing == expression {
			return nil
		}
	}
	ts.tags[label] = append(ts.tags[label], expression)
	return nil
}

// Remove deletes the association between a label and an expression.
func (ts *TagStore) Remove(label, expression string) error {
	exprs, ok := ts.tags[label]
	if !ok {
		return fmt.Errorf("tag %q not found", label)
	}
	updated := exprs[:0]
	for _, e := range exprs {
		if e != expression {
			updated = append(updated, e)
		}
	}
	if len(updated) == 0 {
		delete(ts.tags, label)
	} else {
		ts.tags[label] = updated
	}
	return nil
}

// ByLabel returns all expressions associated with a label, sorted.
func (ts *TagStore) ByLabel(label string) []string {
	exprs := ts.tags[label]
	result := make([]string, len(exprs))
	copy(result, exprs)
	sort.Strings(result)
	return result
}

// Labels returns all known labels in sorted order.
func (ts *TagStore) Labels() []string {
	labels := make([]string, 0, len(ts.tags))
	for l := range ts.tags {
		labels = append(labels, l)
	}
	sort.Strings(labels)
	return labels
}
