package commentstatus

// CommentStatus status
type CommentStatus int

const (
	// Deleted enquiry deleted
	Deleted CommentStatus = iota
	// Viewed enquiry viewed
	Viewed
	// NotViewed enquiry not viewed
	NotViewed
)