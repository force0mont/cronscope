package cron

import (
	"strings"
	"time"
)

// SearchResult holds a matched cron expression and its next occurrence.
type SearchResult struct {
	Expression string
	Description string
	NextRun time.Time
}

// Search filters a list of cron expressions by a query string,
// matching against the expression text and optional description.
// It returns up to maxResults results with their next scheduled run.
func Search(expressions []string, query string, maxResults int, from time.Time) []SearchResult {
	if maxResults <= 0 {
		maxResults = 10
	}
	query = strings.ToLower(strings.TrimSpace(query))
	var results []SearchResult

	for _, expr := range expressions {
		if len(results) >= maxResults {
			break
		}
		normalized := strings.ToLower(expr)
		if query == "" || strings.Contains(normalized, query) {
			sched, err := Parse(expr)
			if err != nil {
				continue
			}
			next := sched.Next(from)
			results = append(results, SearchResult{
				Expression: expr,
				NextRun: next,
			})
		}
	}
	return results
}

// SearchBookmarks filters bookmark expressions by query.
func SearchBookmarks(store *BookmarkStore, query string, maxResults int, from time.Time) []SearchResult {
	bookmarks := store.All()
	exprs := make([]string, 0, len(bookmarks))
	for _, b := range bookmarks {
		exprs = append(exprs, b.Expression)
	}
	return Search(exprs, query, maxResults, from)
}
