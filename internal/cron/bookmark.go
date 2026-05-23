package cron

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

// Bookmark represents a saved cron expression with a user-defined label.
type Bookmark struct {
	Label     string    `json:"label"`
	Expr      string    `json:"expr"`
	Timezone  string    `json:"timezone"`
	CreatedAt time.Time `json:"created_at"`
}

// BookmarkStore manages persisted bookmarks.
type BookmarkStore struct {
	path  string
	items []Bookmark
}

// NewBookmarkStore creates a store backed by the given file path.
func NewBookmarkStore(path string) (*BookmarkStore, error) {
	s := &BookmarkStore{path: path}
	if err := s.load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	return s, nil
}

// Add inserts or updates a bookmark by label.
func (s *BookmarkStore) Add(b Bookmark) {
	for i, existing := range s.items {
		if existing.Label == b.Label {
			s.items[i] = b
			return
		}
	}
	s.items = append(s.items, b)
}

// Remove deletes a bookmark by label. Returns false if not found.
func (s *BookmarkStore) Remove(label string) bool {
	for i, b := range s.items {
		if b.Label == label {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return true
		}
	}
	return false
}

// All returns a copy of all bookmarks.
func (s *BookmarkStore) All() []Bookmark {
	out := make([]Bookmark, len(s.items))
	copy(out, s.items)
	return out
}

// Save persists bookmarks to disk.
func (s *BookmarkStore) Save() error {
	data, err := json.MarshalIndent(s.items, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0o644)
}

func (s *BookmarkStore) load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.items)
}
