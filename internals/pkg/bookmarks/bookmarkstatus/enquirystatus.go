package bookmarks

// EnquiryStatus status
type EnquiryStatus int

const (
	// Deleted enquiry deleted
	Deleted EnquiryStatus = iota
	// Viewed enquiry viewed
	Viewed
	// NotViewed enquiry not viewed
	NotViewed
)
