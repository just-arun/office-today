package posts

import (
	"time"

	"github.com/just-arun/office-today/internals/pkg/posts/poststatus"

	"github.com/just-arun/office-today/internals/pkg/comments"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LikePostStruct post
type LikePostStruct struct {
	UserID primitive.ObjectID `json:"userId"`
}

// Tags for tagging posts
type Tags struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

// Posts model
type Posts struct {
	ID          primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	Title       string                `json:"title" bson:"title"`
	Description string                `json:"description" bson:"description"`
	ImageURL    string                `json:"imageUrl" bson:"image_url"`
	UserID      primitive.ObjectID    `json:"userId" bson:"user_id"`
	CommentsID  []primitive.ObjectID  `json:"commentsId" bson:"comments_id"`
	Comments    []comments.Comments   `json:"comments,omitempty" bson:"comments"`
	EnquiryID   []primitive.ObjectID  `json:"enquiryId" bson:"enquiry_id"`
	Likes       []primitive.ObjectID  `json:"likes" bson:"likes"`
	Tags        primitive.ObjectID    `json:"tags" bson:"tags"`
	Status      poststatus.PostStatus `json:"status" bson:"status"`
	CreatedAt   time.Time             `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time             `json:"updatedAt" bson:"updated_at"`
}

// GetPostStruct struct
type GetPostStruct struct {
	ID           primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	Title        string                `json:"title,omitempty" bson:"title"`
	Description  string                `json:"description,omitempty" bson:"description"`
	ImageURL     string                `json:"imageUrl,omitempty" bson:"image_url"`
	Owner        owner                 `json:"owner,omitempty,omitempty" bson:"owner,omitempty"`
	CommentsID   []primitive.ObjectID  `json:"commentsId,omitempty" bson:"comments_id,omitempty"`
	CommentCount int                   `json:"commentCount,omitempty" bson:"comment_count"`
	EnquiryID    []primitive.ObjectID  `json:"enquiryId,omitempty" bson:"enquiry_id"`
	EnquiryCount int                   `json:"enquiryCount" bson:"enquiry_count"`
	Likes        []primitive.ObjectID  `json:"likes" bson:"likes"`
	Tags         primitive.ObjectID    `json:"tags,omitempty" bson:"tags,omitempty"`
	LikeCount    int                   `json:"likeCount" bson:"like_count"`
	Comments     []comments.Comments   `json:"comments,omitempty" bson:"comments"`
	Status       poststatus.PostStatus `json:"status,omitempty" bson:"status"`
	CreatedAt    time.Time             `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt    time.Time             `json:"updatedAt,omitempty" bson:"updated_at"`
}

type owner struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"userName,omitempty" bson:"user_name"`
	ImageURL string             `json:"imageUrl,omitempty" bson:"image_url"`
}

// CommentType comment user type
type CommentType struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Owner     *owner             `json:"owner,omitempty" bson:"owner"`
	Comment   string             `json:"comment,omitempty" bson:"comment"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updated_at"`
}

// User reference for post
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"userName,omitempty" bson:"user_name"`
	Email    string             `json:"email,omitempty" bson:"email"`
	ImageURL string             `json:"imageUrl,omitempty" bson:"image_url"`
}
