package ownerarea


// OwnerArea check owner accesss for a area
type OwnerArea int


const (
	// Post access
	Post OwnerArea = iota
	// Comment access
	Comment
	// Like access
	Like
	// Bookmark access
	Bookmark
	// User access
	User
)