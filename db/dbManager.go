/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client
var DB *mongo.Database

func ConnectDB(ctx context.Context, uri string, dbName string) {
	log.Printf("[DB] Connecting `%s` (dbName : %s)", uri, dbName)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("[DB] MongoDB Initialize Failed :", err)
	}

	pingErr := client.Ping(ctx, nil)

	if pingErr != nil {
		log.Fatal("[DB] MongoDB Doesn't respond to ping :", pingErr)
	}

	Client = client
	DB = client.Database(dbName)

	result, err := DB.ListCollectionNames(
		ctx,
		bson.D{})

	if err != nil {
		log.Fatal("[DB] Fetch MongoDB Collection List Failed!", err)
	}

	log.Println("[DB] Success to connect MongoDB :", DB.Name(), ":", len(result), " Collections found.")

}
