package cron

import (
	"testing"
	"time"
)

var testFrom = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

func TestSearch_EmptyQuery_ReturnsAll(t *testing.T) {
	exprs := []string{"* * * * *", "0 9 * * 1", "30 6 * * *"}
	results := Search(exprs, "", 10, testFrom)
	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}
}

func TestSearch_QueryFilters(t *testing.T) {
	exprs := []string{"0 9 * * 1", "0 9 * * 2", "30 6 * * *"}
	results := Search(exprs, "0 9", 10, testFrom)
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
}

func TestSearch_RespectsMaxResults(t *testing.T) {
	exprs := []string{"* * * * *", "0 9 * * 1", "30 6 * * *", "0 0 * * *"}
	results := Search(exprs, "", 2, testFrom)
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
}

func TestSearch_InvalidExpressionSkipped(t *testing.T) {
	exprs := []string{"not valid", "* * * * *"}
	results := Search(exprs, "", 10, testFrom)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestSearch_NextRunIsAfterFrom(t *testing.T) {
	exprs := []string{"* * * * *"}
	results := Search(exprs, "", 10, testFrom)
	if len(results) == 0 {
		t.Fatal("expected at least one result")
	}
	if !results[0].NextRun.After(testFrom) {
		t.Errorf("expected NextRun after from, got %v", results[0].NextRun)
	}
}

func TestSearch_NoMatch_ReturnsEmpty(t *testing.T) {
	exprs := []string{"* * * * *", "0 9 * * 1"}
	results := Search(exprs, "zzznomatch", 10, testFrom)
	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestSearchBookmarks_ReturnsMatchingBookmarks(t *testing.T) {
	path := tempBookmarkPath(t)
	store := NewBookmarkStore(path)
	_ = store.Add(Bookmark{Expression: "0 9 * * 1", Label: "weekly"})
	_ = store.Add(Bookmark{Expression: "* * * * *", Label: "every minute"})

	results := SearchBookmarks(store, "0 9", 10, testFrom)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Expression != "0 9 * * 1" {
		t.Errorf("unexpected expression: %s", results[0].Expression)
	}
}
