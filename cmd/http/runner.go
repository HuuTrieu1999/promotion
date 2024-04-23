package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	redis2 "github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	http2 "promotion/internal/controller/http"
	"promotion/internal/controller/http/server/http"
	"promotion/internal/core/config"
	"promotion/internal/core/service/campaign"
	"promotion/internal/core/service/voucher"
	"promotion/internal/infra/redis"
	"promotion/internal/infra/repository"
	"syscall"
)

func main() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./conf")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		os.Exit(1)
	}
	instance := gin.New()
	instance.Use(gin.Logger())
	instance.Use(gin.Recovery())
	clientOptions := options.Client().ApplyURI(viper.GetString("mongodb.uri"))
	client, _ := mongo.Connect(context.Background(), clientOptions)
	voucherCollection := client.Database(viper.GetString("mongodb.db-name")).Collection(viper.GetString("mongodb.voucher-collection"))
	campaignCollection := client.Database(viper.GetString("mongodb.db-name")).Collection(viper.GetString("mongodb.campaign-collection"))

	// Initialize repository
	voucherRepo := repository.NewVoucherRepo(voucherCollection)
	campaignRepo := repository.NewCampaignRepo(campaignCollection)

	redisClient := redis2.NewClient(&redis2.Options{
		Addr:     viper.GetString("redis.uri"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	campaignCounter := redis.NewCampaignCounter(redisClient)

	// Create the UserService
	voucherService := voucher.NewVoucherService(voucherRepo)
	campaignService := campaign.NewCampaignService(campaignRepo, campaignCounter, voucherService, viper.GetInt64("voucher.quota"))

	// Create the CampaignController
	campaignController := http2.NewCampaignController(instance, campaignService)

	// Initialize the routes for CampaignController
	campaignController.InitRouter()

	// Create the HTTP server
	httpServer := http.NewHttpServer(
		instance,
		config.HttpServerConfig{
			Port: viper.GetUint("server.http.port"),
		},
	)

	// Start the HTTP server
	httpServer.Start()
	defer func(httpServer http.HttpServer) {
		err := httpServer.Close()
		if err != nil {
			log.Printf("failed to close http server %v", err)
		}
	}(httpServer)

	// Listen for OS signals to perform a graceful shutdown
	log.Println("listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
}
