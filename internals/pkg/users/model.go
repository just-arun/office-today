package users

import (
	"time"

	"github.com/just-arun/office-today/internals/pkg/users/userstatus"
	"github.com/just-arun/office-today/internals/pkg/users/usertype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Bookmark for adding removing bookmark
type Bookmark struct {
	ID primitive.ObjectID `json:"id",bson:"id"`
}

// Users model
type Users struct {
	ID        primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	UserName  string               `json:"userName,omitempty" bson:"user_name"`
	Email     string               `json:"email,omitempty" bson:"email"`
	Password  string               `json:"password,omitempty" bson:"password"`
	Posts     []primitive.ObjectID `json:"posts,omitempty" bson:"posts"`
	Comments  []primitive.ObjectID `json:"comments,omitempty" bson:"comments"`
	Likes     []primitive.ObjectID `json:"likes,omitempty" bson:"likes"`
	Bookmarks []primitive.ObjectID `json:"bookmarks,omitempty" bson:"bookmarks"`
	OTP       int                  `json:"otp,omitempty" bson:"otp,omitempty"`
	ImageURL  string               `json:"imageUrl,omitempty" bson:"image_url"`

	RegistrationNumber       string    `json:"registrationNumber,omitempty" bson:"registration_number"`
	Address                  string    `json:"address,omitempty" bson:"address"`
	PoBox                    string    `json:"poBox,omitempty" bson:"po_box"`
	Phone                    string    `json:"phone,omitempty" bson:"phone"`
	Fax                      string    `json:"fax,omitempty" bson:"fax"`
	Mobile                   string    `json:"mobile,omitempty" bson:"mobile"`
	RegistrationDate         time.Time `json:"registrationDate,omitempty" bson:"registration_date"`
	Subscription             string    `json:"subscription,omitempty" bson:"subscription"`
	PaymentTerms             string    `json:"paymentTerms,omitempty" bson:"payment_terms"`
	ContactPerson            string    `json:"contactPerson,omitempty" bson:"contact_person"`
	ContactPersonDestination string    `json:"contactPersonDescription,omitempty" bson:"contact_person_destination"`

	RefreshToken string                `json:"refreshToken,omitempty" bson:"refresh_token,omitempty"`
	Status       userstatus.UserStatus `json:"userStatus,omitempty" bson:"status"`
	Type         usertype.UserType     `json:"userType,omitempty" bson:"type"`
	CreatedAt    time.Time             `json:"cretedAt" bson:"created_at"`
	UpdatedAt    time.Time             `json:"updatedAt" bson:"updated_at"`
}

// UsersStruct model
type UsersStruct struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id"`
	UserName  string               `json:"userName" bson:"user_name"`
	Email     string               `json:"email" bson:"email"`
	Posts     []primitive.ObjectID `json:"posts" bson:"posts"`
	Comments  []primitive.ObjectID `json:"comments" bson:"comments"`
	Likes     []primitive.ObjectID `json:"likes" bson:"likes"`
	Bookmarks []primitive.ObjectID `json:"bookmarks" bson:"bookmarks"`
	ImageURL  string               `json:"imageUrl" bson:"image_url"`

	RegistrationNumber       string    `json:"registrationNumber" bson:"registration_number"`
	Address                  string    `json:"address" bson:"address"`
	PoBox                    string    `json:"poBox" bson:"po_box"`
	Phone                    string    `json:"phone" bson:"phone"`
	Fax                      string    `json:"fax" bson:"fax"`
	Mobile                   string    `json:"mobile" bson:"mobile"`
	RegistrationDate         time.Time `json:"registrationDate" bson:"registration_date"`
	Subscription             string    `json:"subscription" bson:"subscription"`
	PaymentTerms             string    `json:"paymentTerms" bson:"payment_terms"`
	ContactPerson            string    `json:"contactPerson" bson:"contact_person"`
	ContactPersonDestination string    `json:"contactPersonDescription" bson:"contact_person_destination"`

	Status    userstatus.UserStatus `json:"userStatus" bson:"status"`
	Type      usertype.UserType     `json:"userType" bson:"type"`
	CreatedAt time.Time             `json:"cretedAt" bson:"created_at"`
	UpdatedAt time.Time             `json:"updatedAt" bson:"updated_at"`
}

// SearchStruct for searching user
type SearchStruct struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Email    string             `json:"email" bson:"email"`
	UserName string             `json:"userName" bson:"user_name"`
	Phone    string             `json:"phone" bson:"phone"`
	Mobile   string             `json:"mobile" bson:"mobile"`
	ImageURL  string               `json:"imageUrl" bson:"image_url"`
}
