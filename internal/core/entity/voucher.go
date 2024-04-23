package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type VoucherEntity struct {
	VoucherID   primitive.ObjectID `bson:"_id"`
	UserID      string             `bson:"userID"`
	Discount    int                `bson:"discount"`
	Description string             `bson:"description"`
	ExpireDate  time.Time          `bson:"expireDate"`
}
