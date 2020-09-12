package posts

import (
	"context"
	"fmt"
	"time"

	"github.com/just-arun/office-today/internals/pkg/posts/poststatus"

	"github.com/just-arun/office-today/internals/pkg/comments"

	"go.mongodb.org/mongo-driver/mongo"

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
	p.CommentsID = []primitive.ObjectID{}
	p.Comments = []comments.Comments{}
	p.EnquiryID = []primitive.ObjectID{}
	p.Likes = []primitive.ObjectID{}

	ctx := context.TODO()
	postID, err := collections.
		Post().
		InsertOne(ctx, p)
	if err != nil {
		return "", err
	}

	result, err := collections.
		User().UpdateOne(
		context.TODO(),
		bson.M{"_id": uID},
		bson.M{
			"$push": bson.M{
				"posts": postID.InsertedID,
			},
		},
	)

	if err != nil {
		return "", err
	}

	fmt.Println(result)

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
func GetAll(filter bson.M, page int) ([]GetPostStruct, error) {
	sort := bson.D{{"$sort", bson.M{"created_at": -1}}}
	match := bson.D{{"$match", filter}}
	ownerLookup := bson.D{{"$lookup", bson.M{
		"from":         "users",
		"localField":   "user_id",
		"foreignField": "_id",
		"as":           "owner",
	}}}
	projectData := bson.D{{"$project", bson.M{
		"_id":         1,
		"title":       1,
		"description": 1,
		"image_url":   1,
		"owner":       1,
		"comment_count": bson.M{
			"$size": "$comments_id",
		},
		"enquiry_count": bson.M{
			"$size": "$enquiry_id",
		},
		"likes": 1,
		"like_count": bson.M{
			"$size": "$likes",
		},
		"comments":   1,
		"status":     1,
		"created_at": 1,
		"updated_at": 1,
	}}}
	unwrapOwner := bson.D{{"$unwind", "$owner"}}
	var skip bson.D
	if page > 0 {
		skip = bson.D{{"$skip", (page * 20) - 20}}
	} else {
		skip = bson.D{{"$skip", 0}}
	}

	limit := bson.D{{"$limit", 20}}

	aggregateFilter := mongo.Pipeline{
		match,
		sort,
		ownerLookup,
		unwrapOwner,
		projectData,
		skip,
		limit,
	}

	cursor, err := collections.
		Post().
		Aggregate(context.TODO(), aggregateFilter)

	if err != nil {
		return nil, err
	}

	var posts []GetPostStruct

	for cursor.Next(context.TODO()) {
		var post GetPostStruct
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// GetPostComments get all posts comments
func GetPostComments(postID string) ([]CommentType, error) {
	ID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, err
	}

	matchStage := bson.D{{"$match", bson.D{{"post", ID}}}}
	lookupUser := bson.D{{"$lookup", bson.M{
		"from":         "users",
		"localField":   "user",
		"foreignField": "_id",
		"as":           "owner",
	}}}
	unwindOwner := bson.D{{"$unwind", "$owner"}}
	projectData := bson.D{{"$project", bson.M{
		"_id":        1,
		"owner":      1,
		"comment":    1,
		"created_at": 1,
		"updated_at": 1,
	}}}
	filter := mongo.Pipeline{matchStage, lookupUser, unwindOwner, projectData}
	cursor, err := collections.Comment().Aggregate(
		context.TODO(),
		filter,
	)
	if err != nil {
		return nil, err
	}

	var comment []CommentType

	for cursor.Next(context.TODO()) {
		var com CommentType
		cursor.Decode(&com)
		comment = append(comment, com)
	}
	fmt.Println(comment)
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
		fmt.Println("[err]", err.Error())
		return false
	}
	return true
}

// DeleteOne for deleting one post
func DeleteOne(postID string) (int64, error) {
	pID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return 0, err
	}

	result, err := collections.
		Post().
		DeleteOne(
			context.TODO(),
			bson.M{
				"_id": pID,
			},
		)

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

// EditPost for editing post
func (p *EditPostDto) EditPost(postID string) (*Posts, error) {
	pID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, err
	}
	var post Posts
	result := collections.
		Post().
		FindOneAndUpdate(
			context.TODO(),
			bson.M{
				"_id": pID,
			},
			bson.M{
				"$set": p,
			},
		)
	if err = result.Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

// AddLikeService for like post
func AddLikeService(postID string, userID string) error {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	pID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}
	_, err = collections.Post().UpdateOne(
		context.TODO(),
		bson.M{
			"_id": pID,
		},
		bson.M{
			"$addToSet": bson.M{
				"likes": uID,
			},
		},
	)

	if err != nil {
		return err
	}

	_, err = collections.User().UpdateOne(
		context.TODO(),
		bson.M{
			"_id": uID,
		},
		bson.M{
			"$addToSet": bson.M{
				"likes": pID,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

// RemoveLikeService for like post
func RemoveLikeService(postID string, userID string) error {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	pID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}
	_, err = collections.Post().UpdateOne(
		context.TODO(),
		bson.M{
			"_id": pID,
		},
		bson.M{
			"$pull": bson.M{
				"likes": uID,
			},
		},
	)

	if err != nil {
		return err
	}

	_, err = collections.User().UpdateOne(
		context.TODO(),
		bson.M{
			"_id": uID,
		},
		bson.M{
			"$pull": bson.M{
				"likes": pID,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}
