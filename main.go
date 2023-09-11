package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kavindu-prabhashitha/go-crud-gin-mongodb/config"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {

	var (
		server      *gin.Engine
		ctx         context.Context
		mongoclient *mongo.Client
		redisclient *redis.Client
	)

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Counld not load environment bariables", err)
	}

	// ? Create a context
	ctx = context.TODO()

	// ? Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// ? Connect to Redis
	redisclient = redis.NewClient(&redis.Options{
		Addr: config.RedisUri,
	})
	if _, err := redisclient.Ping(ctx).Result(); err != nil {
		panic(err)
	}
	err = redisclient.Set(ctx, "test", "Welcome to Golang with Redis and MongoDB",
		0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis client connected successfully...")

	// ? Create the Gin Engine instance
	server = gin.Default()

	fmt.Print(config)
}

func main() {
	fmt.Print("This is the main.go file")
}
