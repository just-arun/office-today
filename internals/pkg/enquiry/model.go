package enquiry

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Enquiry structure
type Enquiry struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"userId" bson:"user_id"`
	PostID    primitive.ObjectID `json:"postId" bson:"post_id"`
  Comment   string             `json:"comment" bson:"comment"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at"`
}
