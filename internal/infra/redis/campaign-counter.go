package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"promotion/internal/core/port/service/counter"
)

type campaignCounter struct {
	redisClient *redis.Client
}

func (c campaignCounter) GetAndIncrease(campaignID string) (int64, error) {
	return c.redisClient.Incr(context.Background(),
		fmt.Sprintf("campaign_counter_%s", campaignID)).Result()
}

func NewCampaignCounter(redisClient *redis.Client) counter.CampaignCounter {
	return &campaignCounter{
		redisClient: redisClient,
	}
}
