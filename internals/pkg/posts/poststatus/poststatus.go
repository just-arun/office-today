package poststatus

// PostStatus post status
type PostStatus int

const (
	// Deleted post deleted
	Deleted PostStatus = iota
	// Viewed post viewed
	Viewed
	// NotViewed post not viewed
	NotViewed
)