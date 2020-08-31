package users

import (
	"context"
	"errors"
	"time"

	"github.com/just-arun/office-today/internals/pkg/comments"

	"github.com/just-arun/office-today/internals/pkg/posts"

	Userstatus "github.com/just-arun/office-today/internals/pkg/users/userstatus"

	Usertype "github.com/just-arun/office-today/internals/pkg/users/usertype"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/just-arun/office-today/internals/boot/collections"
)

// Save user
func (u *Users) Save() (string, error) {
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	ctx := context.TODO()
	user, err := collections.User().
		InsertOne(ctx, u)
	if err != nil {
		return "", err
	}
	return user.
		InsertedID.(primitive.ObjectID).
		Hex(), nil
}

// Update user
func Update(
	filter map[string]interface{},
	payload map[string]interface{},
) (interface{}, error) {
	payload["updated_at"] = time.Now().UTC()
	ctx := context.TODO()
	result, err := collections.User().
		UpdateOne(ctx, filter, bson.M{
			"$set": payload,
		})
	if err != nil {
		return nil, err
	}
	return result.UpsertedID, err
}


// CreateAudience create audience by admin
func (u *Users) CreateAudience() (map[string]interface{}, error) {

	dbUser, _ := GetOne(
		map[string]interface{}{
			"email": u.Email,
		},
	)

	if dbUser != nil {
		return nil, errors.New("User already exist")
	}

	u.UserType = Usertype.Publisher
	u.UserStatus = Userstatus.Active

	// saving users
	userID, err := u.Save()

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id": userID,
	}, nil
}

// GetOne user
func GetOne(
	filter map[string]interface{},
) (*Users, error) {
	var user Users
	ctx := context.TODO()
	err := collections.User().
		FindOne(ctx, filter).
		Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAll users
func GetAll(
	filter map[string]interface{},
	page int,
) ([]*Users, error) {
	var users []*Users

	option := options.Find()
	skip := int64((page * 20) - 20)
	limit := int64(20)

	option.Skip = &skip
	option.Limit = &limit
	option.Sort = bson.D{{"createdAt", -1}}

	ctx := context.TODO()
	cursor, err := collections.User().
		Find(ctx, filter, option)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user *Users
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserPosts get all user posts
func GetUserPosts(userID string, page int) ([]*posts.Posts, error) {

	ID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return nil, err
	}

	userPosts, err := posts.GetAllPost(
		bson.M{
			"user_id": ID,
		},
		page,
	)

	if err != nil {
		return nil, err
	}

	return userPosts, nil
}

// GetUserComments get all users posts
func GetUserComments(userID string) ([]*comments.Comments, error) {
	ID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return nil, err
	}

	comment, err := comments.GetAllComments(bson.M{"user_id": ID})

	if err != nil {
		return nil, err
	}

	return comment, nil
}

// UpdateRefreshToken update refresh token
func UpdateRefreshToken(userID string, rToken string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = collections.User().UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"refresh_token": rToken,
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}