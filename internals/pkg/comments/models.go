package comments

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comments module for posts
type Comments struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user" bson:"user"`
	PostID    primitive.ObjectID `json:"post" bson:"post"`
	Comment   string             `json:"comment" bson:"comment"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at"`
}
