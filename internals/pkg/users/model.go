package users

import (
	"time"

	"github.com/just-arun/office-today/internals/pkg/comments"

	"github.com/just-arun/office-today/internals/pkg/posts"

	"github.com/just-arun/office-today/internals/pkg/users/userstatus"
	"github.com/just-arun/office-today/internals/pkg/users/usertype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Users model
type Users struct {
	ID         primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	UserName   string                `json:"userName" bson:"user_name"`
	Email      string                `json:"email" bson:"email"`
	Password   string                `json:"password" bson:"password"`
	PostID     []primitive.ObjectID  `json:"PostId" bson:"post_id"`
	Posts      []*posts.Posts        `json:"Posts,omitempty" bson:"posts"`
	CommentID  []primitive.ObjectID  `json:"CommentId" bson:"comment_id"`
	Comments   []*comments.Comments  `json:"comments,omitempty" bson:"comments"`
	Likes      []primitive.ObjectID  `json:"likes" bson:"likes,omitempty"`
	UserStatus userstatus.UserStatus `json:"userStatus" bson:"user_status,omitempty"`
	UserType   usertype.UserType     `json:"userType" bson:"user_type,omitempty"`
	CreatedAt  time.Time             `json:"cretedAt" bson:"created_at,omitempty"`
	UpdatedAt  time.Time             `json:"updatedAt" bson:"updated_at,omitempty"`
}
