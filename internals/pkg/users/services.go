package users

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/just-arun/office-today/internals/util/message"
	"github.com/just-arun/office-today/internals/util/password"

	"github.com/just-arun/office-today/internals/pkg/comments"
	"github.com/just-arun/office-today/internals/pkg/posts/poststatus"

	"github.com/just-arun/office-today/internals/pkg/posts"

	Userstatus "github.com/just-arun/office-today/internals/pkg/users/userstatus"
	"github.com/just-arun/office-today/internals/pkg/users/usertype"

	Usertype "github.com/just-arun/office-today/internals/pkg/users/usertype"

	"go.mongodb.org/mongo-driver/mongo"
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

// GetOneUserService for getting user
func GetOneUserService(filter bson.M)  (*UsersStruct, error) {
	var user UsersStruct
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
) ([]*UsersStruct, error) {
	var users []*UsersStruct

	option := options.Find()
	amount := 10

	if page > 0 {
		skip := int64((page * amount) - amount)
		limit := int64(amount)

		option.Skip = &skip
		option.Limit = &limit
	}
	option.Sort = bson.M{"createdAt": -1}
	option.Projection = bson.M{
		"_id":                        1,
		"user_name":                  1,
		"email":                      1,
		"posts":                      1,
		"comments":                   1,
		"likes":                      1,
		"bookmarks":                  1,
		"otp":                        1,
		"image_url":                  1,
		"registration_number":        1,
		"address":                    1,
		"po_box":                     1,
		"phone":                      1,
		"fax":                        1,
		"mobile":                     1,
		"registration_date":          1,
		"subscription":               1,
		"payment_terms":              1,
		"contact_person":             1,
		"contact_person_destination": 1,
		"status":                     1,
		"type":                       1,
		"created_at":                 1,
		"updated_at":                 1,
	}

	ctx := context.TODO()
	cursor, err := collections.User().
		Find(ctx, filter, option)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user *UsersStruct
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
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
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	pwd, err := password.Encrypt(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = pwd
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

// GetUserProfileService for getting user profile
func GetUserProfileService(userID string) (*UsersStruct, error) {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	match := bson.D{{
		Key: "$match",
		Value: bson.M{
			"_id": uID,
			"status": bson.M{
				"$ne": Userstatus.Disabled,
			},
		},
	}}
	project := bson.D{{
		Key: "$project",
		Value: bson.M{
			"_id":                        1,
			"user_name":                  1,
			"email":                      1,
			"posts":                      1,
			"comments":                   1,
			"likes":                      1,
			"bookmarks":                  1,
			"image_url":                  1,
			"registration_number":        1,
			"address":                    1,
			"po_box":                     1,
			"phone":                      1,
			"fax":                        1,
			"mobile":                     1,
			"registration_date":          1,
			"subscription":               1,
			"payment_terms":              1,
			"contact_person":             1,
			"contact_person_destination": 1,
			"type":                       1,
			"created_at":                 1,
			"updated_at":                 1,
		},
	}}
	filter := mongo.Pipeline{match, project}
	cursor, err := collections.
		User().
		Aggregate(
			context.TODO(),
			filter,
		)

	if err != nil {
		return nil, err
	}
	var user UsersStruct

	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println(user)

	return &user, nil
}

// SearchUserService for searching users
func SearchUserService(key string) ([]SearchStruct, error) {

	mod := mongo.IndexModel{
		Keys: bson.D{
			{Key: "$**", Value: "text"},
		},
		Options: &options.IndexOptions{},
	}
	ind, err := collections.User().Indexes().CreateOne(context.TODO(), mod)
	if err != nil {
		return nil, err
	}
	fmt.Println("index", ind)

	cursor, err := collections.User().Find(context.TODO(), bson.M{
		"type": bson.M{
			"$ne": usertype.Admin,
		},
		"$text": bson.M{
			"$search": key,
		},
	})
	if err != nil {
		return nil, err
	}
	var users []SearchStruct
	for cursor.Next(context.TODO()) {
		var user SearchStruct
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserPostServices for getting all posts
func GetUserPostServices(page int, userID string) ([]*posts.GetPostStruct, error) {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{Key: "$match", Value: bson.M{
			"user_id": uID,
			"status": bson.M{
				"$ne": poststatus.Deleted,
			},
		}},
	}
	perPage := 20
	var skip bson.D
	if page > 0 {
		skip = bson.D{{Key: "$skip", Value: (page * perPage) - perPage}}
	} else {
		skip = bson.D{{Key: "$skip", Value: 0}}
	}

	limit := bson.D{{Key: "$limit", Value: perPage}}

	cursor, err := collections.Post().Aggregate(context.TODO(),
		mongo.Pipeline{
			filter,
			skip,
			limit,
		},
	)

	if err != nil {
		return nil, err
	}
	var postsList []*posts.GetPostStruct
	for cursor.Next(context.TODO()) {
		var post *posts.GetPostStruct
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		postsList = append(postsList, post)
	}

	if postsList == nil {
		postsList = []*posts.GetPostStruct{}
	}

	return postsList, nil
}

// GetUserFavServices for getting all posts
func GetUserFavServices(page int, userID string) ([]*posts.GetPostStruct, error) {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	var user Users
	if err := collections.User().
		FindOne(context.TODO(),
			bson.M{
				"_id": uID,
			}).Decode(&user); err != nil {
		return nil, err
	}

	fmt.Printf("\n[Linkes]: %v\n", user.Likes);

	match := []bson.M{}
	
	match = append(match, bson.M{
		"status": bson.M{
			"$ne": poststatus.Deleted,
		},
	})
	match = append(match, bson.M{
		"_id": bson.M{
			"$in": user.Likes,
		},
	})
	
	fmt.Println(match)

	filter := bson.D{
		{Key: "$match", Value: bson.M{
			"$and": match,
		}},
	}

	perPage := 20
	var skip bson.D
	if page > 0 {
		skip = bson.D{{Key: "$skip", Value: (page * perPage) - perPage}}
	} else {
		skip = bson.D{{Key: "$skip", Value: 0}}
	}

	limit := bson.D{{Key: "$limit", Value: perPage}}

	cursor, err := collections.Post().Aggregate(context.TODO(),
		mongo.Pipeline{
			filter,
			skip,
			limit,
		},
	)

	if err != nil {
		return nil, err
	}
	var postsList []*posts.GetPostStruct
	for cursor.Next(context.TODO()) {
		var post *posts.GetPostStruct
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		postsList = append(postsList, post)
	}
	fmt.Printf("\nlikes%v\n", postsList)

	if postsList == nil {
		postsList = []*posts.GetPostStruct{}
	}

	return postsList, nil
}


// ResetPassword for reseting password
func ResetPassword(email string) (error) {
	
	err := message.Mail(email, "forgot password", "this is test message");
	// collections.User().FindOneAndUpdate()
	return err;
}