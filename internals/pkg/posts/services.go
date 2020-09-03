package posts

import (
	"context"
	"fmt"
	"time"

	"github.com/just-arun/office-today/internals/pkg/posts/poststatus"

	"github.com/just-arun/office-today/internals/pkg/comments"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/just-arun/office-today/internals/boot/collections"
)

// Save posts
func (p *Posts) Save(userID string) (string, error) {

	uID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return "", err
	}

	p.Status = poststatus.NotViewed
	p.UserID = uID
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

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

// GetAll get all posts
func GetAll(filter bson.M, page int) ([]*Posts, error) {
	var posts []*Posts

	option := options.Find()
	count := 20
	skip := int64((page * count) - count)
	limit := int64(count * 1)

	if page > 0 {
		option.Skip = &skip
		option.Limit = &limit
	}
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

// CheckOwner for post
func CheckOwner(postID string, userID string) bool {
	pID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	var post Posts
	if err := collections.
		Post().
		FindOne(
			context.TODO(),
			bson.M{
				"_id":     pID,
				"user_id": uID,
			},
		).
		Decode(&post); err != nil {
		fmt.Println(err)
		return false
	}
	return true
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
