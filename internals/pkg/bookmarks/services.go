package bookmarks

import (
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"github.com/just-arun/office-today/internals/boot/collections"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CheckOwner for bookmark
func CheckOwner(bookmarkID string, userID string) bool {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return false
	}
	bID, err := primitive.ObjectIDFromHex(bookmarkID)
	if err != nil {
		return false
	}
	var bookmark Bookmark
	if err := collections.
		Bookmarks().
		FindOne(
			context.TODO(),
			bson.M{
				"_id": bID,
				"user_id": uID,
			},
		).
		Decode(&bookmark);
		err != nil {
		return false
	}
	return true
}