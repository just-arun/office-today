package posts

import "go.mongodb.org/mongo-driver/bson/primitive"

// CreatePostDto used for creating post
type CreatePostDto struct {
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	ImageURL    string             `json:"imageUrl" bson:"image_url"`
	Tags        primitive.ObjectID `json:"tags" bson:"tags"`
}

// EditPostDto for editing post
type EditPostDto struct {
	Title       string `json:"title,omitempty" bson:"title,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	ImageURL    string `json:"imageUrl,omitempty" bson:"image_url,omitempty"`
}
