package users

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/just-arun/office-today/internals/boot/collections"
)

// Save user
func (u *Users) Save() (interface{}, error) {
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	ctx := context.TODO()
	user, err := collections.User().
		InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}
  return user.
    InsertedID.
    (primitive.ObjectID).
    Hex, nil
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

// GetOne user
func GetOne(
    filter map[string]interface{},
    option options.FindOneOptions,
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
	option *options.FindOptions,
) ([]*Users, error) {
	var users []*Users
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


