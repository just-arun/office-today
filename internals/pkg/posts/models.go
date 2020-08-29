package posts

import (
	"time"

	"github.com/just-arun/office-today/internals/pkg/posts/poststatus"

	"github.com/just-arun/office-today/internals/pkg/comments"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Posts model
type Posts struct {
	ID          primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	Description string                `json:"description" bson:"description"`
	ImageURL    string                `json:"imageUrl" bson:"image_url"`
	UserID      primitive.ObjectID    `json:"userId" bson:"user_id"`
	CommentsID  []primitive.ObjectID  `json:"commentsId" bson:"comments_id"`
	Comments    []*comments.Comments  `json:"comments" bson:"comments"`
	EnquiryID   []primitive.ObjectID  `json:"enquiryId" bson:"enquiry_id"`
	Likes       []primitive.ObjectID  `json:"likes" bson:"likes"`
	Status      poststatus.PostStatus `json:"status" bson:"status"`
	CreatedAt   time.Time             `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time             `json:"updatedAt" bson:"updated_at"`
}
