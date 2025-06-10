package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

const (
	// GitHubCopilotChangelogURL is the RSS feed URL for GitHub Copilot changelog
	GitHubCopilotChangelogURL = "https://github.blog/changelog/label/copilot/feed/"
)

// ChangelogReader handles reading and filtering RSS feeds
type ChangelogReader struct {
	parser *gofeed.Parser
}

// NewChangelogReader creates a new ChangelogReader instance
func NewChangelogReader() *ChangelogReader {
	return &ChangelogReader{
		parser: gofeed.NewParser(),
	}
}

// GetTodaysPosts fetches RSS feed and returns only posts published today
func (cr *ChangelogReader) GetTodaysPosts(feedURL string) ([]*gofeed.Item, error) {
	feed, err := cr.parser.ParseURL(feedURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	return cr.filterTodaysPosts(feed), nil
}

// GetTodaysPostsFromString parses RSS from string and returns only posts published today
func (cr *ChangelogReader) GetTodaysPostsFromString(feedXML string) ([]*gofeed.Item, error) {
	feed, err := cr.parser.ParseString(feedXML)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	return cr.filterTodaysPosts(feed), nil
}

// filterTodaysPosts filters feed items to return only today's posts
func (cr *ChangelogReader) filterTodaysPosts(feed *gofeed.Feed) []*gofeed.Item {
	today := time.Now().UTC().Truncate(24 * time.Hour)

	var todaysPosts []*gofeed.Item
	for _, item := range feed.Items {
		if item.PublishedParsed != nil {
			publishedDate := item.PublishedParsed.UTC().Truncate(24 * time.Hour)
			if publishedDate.Equal(today) {
				todaysPosts = append(todaysPosts, item)
			}
		}
	}

	return todaysPosts
}

func main() {
	reader := NewChangelogReader()

	posts, err := reader.GetTodaysPosts(GitHubCopilotChangelogURL)
	if err != nil {
		log.Fatalf("Error reading RSS feed: %v", err)
	}

	fmt.Printf("Found %d posts published today:\n\n", len(posts))

	for i, post := range posts {
		fmt.Printf("%d. %s\n", i+1, post.Title)
		fmt.Printf("   Published: %s\n", post.PublishedParsed.Format("2006-01-02 15:04:05 UTC"))
		fmt.Printf("   Link: %s\n\n", post.Link)
	}
}
