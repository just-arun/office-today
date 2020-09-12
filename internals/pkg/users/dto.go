package users

// UpdateUserStruct for updating user
type UpdateUserStruct struct {
	ImageURL string `json:"imageUrl,omitempty" bson:"image_url,omitempty"`
	UserName string `json:"userName,omitempty" bson:"user_name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
}
