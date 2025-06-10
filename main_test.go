package main

import (
	"strings"
	"testing"
	"time"
)

// mockRSSFeed returns a mock RSS feed XML for testing
func mockRSSFeed() string {
	today := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 MST")
	yesterday := time.Now().UTC().AddDate(0, 0, -1).Format("Mon, 02 Jan 2006 15:04:05 MST")
	
	return `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
<channel>
<title>GitHub Changelog - Copilot</title>
<description>Recent changes to GitHub Copilot</description>
<link>https://github.blog/changelog/label/copilot/</link>
<item>
<title>Today's Copilot Update</title>
<link>https://github.blog/changelog/2023/01/01/todays-copilot-update/</link>
<description>A new feature for Copilot</description>
<pubDate>` + today + `</pubDate>
</item>
<item>
<title>Yesterday's Copilot Update</title>
<link>https://github.blog/changelog/2023/01/01/yesterdays-copilot-update/</link>
<description>An old feature for Copilot</description>
<pubDate>` + yesterday + `</pubDate>
</item>
<item>
<title>Another Today's Update</title>
<link>https://github.blog/changelog/2023/01/01/another-todays-update/</link>
<description>Another new feature for Copilot</description>
<pubDate>` + today + `</pubDate>
</item>
</channel>
</rss>`
}

func TestGetTodaysPosts(t *testing.T) {
	reader := NewChangelogReader()
	
	// Test with mock RSS feed using the new method
	mockFeedXML := mockRSSFeed()
	todaysPosts, err := reader.GetTodaysPostsFromString(mockFeedXML)
	if err != nil {
		t.Fatalf("Failed to get today's posts: %v", err)
	}
	
	// Should have 2 posts from today
	if len(todaysPosts) != 2 {
		t.Errorf("Expected 2 posts from today, got %d", len(todaysPosts))
	}
	
	// Check that both posts are from today
	for _, post := range todaysPosts {
		if !strings.Contains(post.Title, "Today") {
			t.Errorf("Expected post title to contain 'Today', got: %s", post.Title)
		}
	}
}

func TestGetTodaysPostsWithURL(t *testing.T) {
	// This test would require internet access, so we'll use a mock server
	// For now, we'll test the parsing logic with a string
	reader := NewChangelogReader()
	
	mockFeedXML := mockRSSFeed()
	feed, err := reader.parser.ParseString(mockFeedXML)
	if err != nil {
		t.Fatalf("Failed to parse mock RSS feed: %v", err)
	}
	
	// Verify we can parse the feed structure correctly
	if len(feed.Items) != 3 {
		t.Errorf("Expected 3 items in feed, got %d", len(feed.Items))
	}
	
	if feed.Title != "GitHub Changelog - Copilot" {
		t.Errorf("Expected feed title 'GitHub Changelog - Copilot', got: %s", feed.Title)
	}
}

func TestNewChangelogReader(t *testing.T) {
	reader := NewChangelogReader()
	
	if reader == nil {
		t.Error("NewChangelogReader() returned nil")
	}
	
	if reader.parser == nil {
		t.Error("Parser not initialized in ChangelogReader")
	}
}

func TestGetTodaysPostsFromStringError(t *testing.T) {
	reader := NewChangelogReader()
	
	// Test with invalid XML
	_, err := reader.GetTodaysPostsFromString("invalid xml")
	if err == nil {
		t.Error("Expected error for invalid XML, got nil")
	}
}

func TestFilterTodaysPostsEmptyFeed(t *testing.T) {
	reader := NewChangelogReader()
	
	// Test with empty RSS feed
	emptyFeedXML := `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
<channel>
<title>Empty Feed</title>
<description>Empty feed for testing</description>
<link>https://example.com/</link>
</channel>
</rss>`
	
	todaysPosts, err := reader.GetTodaysPostsFromString(emptyFeedXML)
	if err != nil {
		t.Fatalf("Failed to parse empty feed: %v", err)
	}
	
	if len(todaysPosts) != 0 {
		t.Errorf("Expected 0 posts from empty feed, got %d", len(todaysPosts))
	}
}