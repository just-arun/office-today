package comments

import (
	"context"

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
