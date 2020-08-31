package users

import (
	"time"

	"github.com/just-arun/office-today/internals/pkg/users/userstatus"
	"github.com/just-arun/office-today/internals/pkg/users/usertype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Users model
type Users struct {
	ID        primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	UserName  string                `json:"userName,omitempty" bson:"user_name"`
	Email     string                `json:"email,omitempty" bson:"email"`
	Password  string                `json:"password,omitempty" bson:"password"`
	Posts     []*primitive.ObjectID `json:"Posts,omitempty" bson:"posts"`
	Comments  []*primitive.ObjectID `json:"comments,omitempty" bson:"comments"`
	Likes     []*primitive.ObjectID `json:"likes,omitempty" bson:"likes,omitempty"`
	Bookmarks []*primitive.ObjectID `json:"bookmarks,omitempty" bson:"bookmarks,omitempty"`

	RegistrationNumber       string    `json:"registrationNumber,omitempty" bson:"registration_number"`
	Address                  string    `json:"address,omitempty" bson:"address"`
	PoBox                    string    `json:"poBox,omitempty" bson:"po_box"`
	Phone                    int       `json:"phone,omitempty" bson:"phone"`
	Fax                      int       `json:"fax,omitempty" bson:"fax"`
	Mobile                   int       `json:"mobile,omitempty" bson:"mobile"`
	RegistrationDate         time.Time `json:"registrationDate,omitempty" bson:"registration_date"`
	Subscription             string    `json:"subscription,omitempty" bson:"subscription"`
	PaymentTerms             string    `json:"paymentTerms,omitempty" bson:"payment_terms"`
	ContactPerson            string    `json:"contactPerson,omitempty" bson:"contact_person"`
	ContactPersonDestination string    `json:"contactPersonDescription,omitempty" bson:"contact_person_destination"`

	UserStatus userstatus.UserStatus `json:"userStatus,omitempty" bson:"user_status"`
	UserType   usertype.UserType     `json:"userType,omitempty" bson:"user_type"`
	CreatedAt  time.Time             `json:"cretedAt" bson:"created_at"`
	UpdatedAt  time.Time             `json:"updatedAt" bson:"updated_at"`
}
