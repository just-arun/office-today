package comments

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/just-arun/office-today/internals/boot/collections"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllCommentsService all comments for the post
func GetAllCommentsService(filter map[string]interface{}, comments []*Comments) error {
	option := options.Find()
	option.Sort = bson.M{"created_at": -1}
	ctx := context.TODO()

	cursor, err := collections.
		Comment().
		Find(ctx, filter, option)

	if err != nil {
		return err
	}

	for cursor.Next(context.TODO()) {
		var comment *Comments
		err := cursor.Decode(&comment)
		if err != nil {
			return err
		}
		comments = append(comments, comment)
	}

	return nil
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
		Comment().
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

// DeleteCommentService for
func DeleteCommentService(id string) error {

	cID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	commentFilter := bson.M{"_id": cID}

	var comment Comments

	if err := collections.
		Comment().
		FindOneAndDelete(context.TODO(), commentFilter).
		Decode(&comment); err != nil {
		return err
	}

	postFilter := bson.M{"_id": comment.PostID}
	postCommentData := bson.M{
		"$pull": bson.M{
			"comments_id": comment.ID,
		},
	}

	_, err = collections.
		User().
		UpdateOne(
			context.TODO(),
			postFilter,
			postCommentData,
		)

	if err != nil {
		return err
	}

	userFilter := bson.M{"_id": comment.UserID}
	userFilterData := bson.M{
		"$pull": bson.M{
			"comments": comment.ID,
		},
	}

	_, err = collections.
		User().
		UpdateOne(
			context.TODO(),
			userFilter,
			userFilterData,
		)
	if err != nil {
		return err
	}

	return nil
}
