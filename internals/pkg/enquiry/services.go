package enquiry

import (
	"context"
	"fmt"

	"github.com/just-arun/office-today/internals/boot/collections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CheckOwner for enquiry
func CheckOwner(enquiryID string, userID string) bool {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	eID, err := primitive.ObjectIDFromHex(enquiryID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	var enquiry Enquiry
	if err := collections.
		Enquiry().
		FindOne(
			context.TODO(),
			bson.M{
				"_id":     eID,
				"user_id": uID,
			},
		).
		Decode(&enquiry); err != nil {
		return false
	}
	return true
}
