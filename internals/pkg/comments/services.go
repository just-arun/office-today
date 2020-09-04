package comments

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/just-arun/office-today/internals/boot/collections"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllComments all comments for the post
func GetAllComments(filter map[string]interface{}) ([]*Comments, error) {

	var comments []*Comments
	option := options.Find()
	option.Sort = bson.M{"created_at": -1}
	ctx := context.TODO()

	cursor, err := collections.
		Comment().
		Find(ctx, filter, option)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var comment *Comments
		err := cursor.Decode(&comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

// Save for creating comment
func (c *Comments) Save(userID string) (string, error) {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return "", err
	}
	c.UserID = uID
	result, err := collections.
		Comment().
		InsertOne(context.TODO(), c)
	if err != nil {
		return "", err
	}

	fmt.Println(c.PostID)

	_, err = collections.Post().UpdateOne(
		context.TODO(),
		bson.M{
			"_id": c.PostID,
		},
		bson.M{
			"$push": bson.M{
				"comments_id": result.InsertedID,
			},
		})
	if err != nil {
		return "", err
	}

	_, err = collections.User().UpdateOne(
		context.TODO(),
		bson.M{
			"_id": uID,
		},
		bson.M{
			"$push": bson.M{
				"comments": result.InsertedID,
			},
		})
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).
		Hex(), nil
}

// CheckOwner for bookmark
func CheckOwner(commentID string, userID string) bool {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return false
	}
	cID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return false
	}
	var comment Comments
	if err := collections.
		Bookmarks().
		FindOne(
			context.TODO(),
			bson.M{
				"_id":     cID,
				"user_id": uID,
			},
		).
		Decode(&comment); err != nil {
		return false
	}
	return true
}
