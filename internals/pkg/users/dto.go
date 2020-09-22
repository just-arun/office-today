package users

import (
	"time"

	"github.com/just-arun/office-today/internals/pkg/users/userstatus"
	"github.com/just-arun/office-today/internals/pkg/users/usertype"
)

// UpdateUserStruct for updating user
type UpdateUserStruct struct {
	UserName string `json:"userName,omitempty" bson:"user_name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	ImageURL string `json:"imageUrl,omitempty" bson:"image_url,omitempty"`

	RegistrationNumber       string    `json:"registrationNumber,omitempty" bson:"registration_number,omitempty"`
	Address                  string    `json:"address,omitempty" bson:"address,omitempty"`
	PoBox                    string    `json:"poBox,omitempty" bson:"po_box,omitempty"`
	Phone                    string    `json:"phone,omitempty" bson:"phone,omitempty"`
	Fax                      string    `json:"fax,omitempty" bson:"fax,omitempty"`
	Mobile                   string    `json:"mobile,omitempty" bson:"mobile,omitempty"`
	RegistrationDate         time.Time `json:"registrationDate,omitempty" bson:"registration_date,omitempty"`
	Subscription             string    `json:"subscription,omitempty" bson:"subscription,omitempty"`
	PaymentTerms             string    `json:"paymentTerms,omitempty" bson:"payment_terms,omitempty"`
	ContactPerson            string    `json:"contactPerson,omitempty" bson:"contact_person,omitempty"`
	ContactPersonDestination string    `json:"contactPersonDescription,omitempty" bson:"contact_person_destination,omitempty"`

	Status    userstatus.UserStatus `json:"userStatus,omitempty" bson:"status,omitempty"`
	Type      usertype.UserType     `json:"userType,omitempty" bson:"type,omitempty"`
	CreatedAt time.Time             `json:"cretedAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time             `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
}
