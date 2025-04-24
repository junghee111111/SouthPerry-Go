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
	"os"
	"time"
)

var Client *mongo.Client
var DB *mongo.Database

func Init(uri string, dbName string) {
	log.Printf("[DB] %s (dbName : %s) 에 연결합니다.", uri, dbName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("[DB] MongoDB Initialize 실패. 잘못된 주소인 것 같습니다.:", err)
		os.Exit(1)
	}

	pingErr := client.Ping(ctx, nil)

	if pingErr != nil {
		log.Fatal("[DB] MongoDB에 연결할 수 없습니다!:", pingErr)
		os.Exit(1)
	}

	Client = client
	DB = client.Database(dbName)

	result, err := DB.ListCollectionNames(
		ctx,
		bson.D{})

	if err != nil {
		log.Fatal("[DB] 컬렉션 조회 실패!", err)
		os.Exit(1)
	}

	log.Println("[DB] MongoDB 연결 완료:", DB.Name(), ":", len(result), "개의 컬렉션")

}
