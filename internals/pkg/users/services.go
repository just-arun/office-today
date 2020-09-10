package users

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/just-arun/office-today/internals/util/password"

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
	pwd, err := password.Encrypt(u.Password)
	if err != nil {
		return "", err
	}
	u.Password = pwd
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	u.Posts = []primitive.ObjectID{}
	u.Comments = []primitive.ObjectID{}
	u.Likes = []primitive.ObjectID{}
	u.Bookmarks = []primitive.ObjectID{}
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

// UpdateUserService for updating user
func UpdateUserService(userID string, payload bson.M) error {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": uID}
	result, err := collections.
		User().
		UpdateOne(
			context.TODO(),
			filter,
			payload,
		)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

// CreateAudience create audience by admin
func (u *Users) CreateAudience() (map[string]interface{}, error) {

	dbUser, _ := GetOne(
		bson.M{
			"email": u.Email,
		},
	)

	if dbUser != nil {
		return nil, errors.New("User already exist")
	}

	u.Type = Usertype.Publisher
	u.Status = Userstatus.Active

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
	filter bson.M,
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
	if page > 0 {
		skip := int64((page * 20) - 20)
		limit := int64(20)

		option.Skip = &skip
		option.Limit = &limit
	}
	option.Sort = bson.M{"createdAt": -1}

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
func GetUserPosts(userID string, page int) ([]posts.Posts, error) {

	ID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return nil, err
	}

	userPosts, err := posts.GetAll(
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
func GetUserComments(userID string, comment []*comments.Comments) error {
	ID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return err
	}

	err = comments.
		GetAllCommentsService(bson.M{"user_id": ID}, comment)

	if err != nil {
		return err
	}

	return nil
}

// CreateUserService for creating user
func CreateUserService(user Users) (string, error) {
	ID, err := collections.User().InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}
	return ID.InsertedID.(primitive.ObjectID).Hex(), nil
}

// AddBookmarkService for adding bookmark
func AddBookmarkService(userID string, postID primitive.ObjectID) error {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	result, err := collections.
		User().UpdateOne(
		context.TODO(),
		bson.M{"_id": uID}, bson.M{
			"$addToSet": bson.M{
				"bookmarks": postID,
			},
		})
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

// RemoveBookmarkService for adding bookmark
func RemoveBookmarkService(userID string, postID primitive.ObjectID) error {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	result, err := collections.
		User().UpdateOne(
		context.TODO(),
		bson.M{"_id": uID}, bson.M{
			"$pull": bson.M{
				"bookmarks": postID,
			},
		})
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

