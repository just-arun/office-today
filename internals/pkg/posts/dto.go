package posts

// CreatePostDto used for creating post
type CreatePostDto struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	ImageURL    string `json:"imageUrl" bson:"image_url"`
}
