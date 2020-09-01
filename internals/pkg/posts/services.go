package posts

import (
	"context"

	"github.com/just-arun/office-today/internals/pkg/posts/poststatus"

	"github.com/just-arun/office-today/internals/pkg/comments"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/just-arun/office-today/internals/boot/collections"
)

// Save posts
func (p *Posts) Save() (string, error) {

	p.Status = poststatus.NotViewed

	ctx := context.TODO()
	postID, err := collections.
		Post().
		InsertOne(ctx, p)
	if err != nil {
		return "", err
	}

	return postID.
		InsertedID.(primitive.ObjectID).
		Hex(), nil
}

// GetOne post
func GetOne(fileter bson.M) (*Posts, error) {
	var post Posts

	ctx := context.TODO()
	err := collections.
		Post().
		FindOne(ctx, fileter).Decode(&post)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

// GetAllPost get all posts
func GetAllPost(filter bson.M, page int) ([]*Posts, error) {
	var posts []*Posts

	option := options.Find()
	count := 20
	skip := int64((page * count) - count)
	limit := int64(count)

	option.Skip = &skip
	option.Limit = &limit
	option.Sort = bson.M{"created_at": -1}

	ctx := context.TODO()

	cursor, err := collections.
		Post().
		Find(ctx, filter, option)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var post Posts
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

// GetPostComments get all posts comments
func GetPostComments(postID string) ([]*comments.Comments, error) {
	ID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, err
	}

	comment, err := comments.GetAllComments(bson.M{"post_id": ID})
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// AddCommentBookmarkLikeEnquiryID add
// comment bookmark like and enquires
// func AddCommentBookmarkLikeEnquiryID(
//     postID string,
//     update bson.M,
//   ) (
//       interface{},
//       error,
//     ) {
// 	ID, err := primitive.ObjectIDFromHex(postID)
// 	if err != nil {
// 		return nil, err
//   }

//   filter := bson.M{ "_id": ID }

//   result, err := collections.
//     Post().
//     UpdateOne(
//       context.TODO(),
//      filter,

//     )

//   return nil, nil
// }
