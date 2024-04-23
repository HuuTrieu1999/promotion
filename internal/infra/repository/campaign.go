package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"promotion/internal/core/port/repository"
	"time"
)

type campaignRepo struct {
	col *mongo.Collection
}

func (c campaignRepo) Insert(campaignID string, userID string) error {
	data := bson.M{}
	data["campaignID"] = campaignID
	data["userID"] = userID
	data["createdDate"] = time.Now()
	_, err := c.col.InsertOne(context.Background(), data)
	return err
}

func NewCampaignRepo(col *mongo.Collection) repository.CampaignRepository {
	return &campaignRepo{
		col: col,
	}
}
