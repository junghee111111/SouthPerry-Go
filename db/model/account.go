/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Account struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	AccId        int                `bson:"accId"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"passwordHash"`
	Sex          bool               `bson:"sex"`
	Birthday     time.Time          `bson:"birthday"`
	IsBanned     bool               `bson:"isBanned"`
	IsLoggedIn   bool               `bson:"isLoggedIn"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
	LastLoggedAt time.Time          `bson:"lastLoggedAt"`
	LastLoggedIp string             `bson:"lastLoggedIp"`
}
