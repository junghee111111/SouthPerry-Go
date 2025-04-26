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
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"passwordHash"`
	AccId        int32              `bson:"accId"`
	IsGM         bool               `bson:"isGM"`
	IsAdult      bool               `bson:"isAdult"`
	Gender       int32              `bson:"gender"`

	CreatedAt time.Time `bson:"createdAt"`
	LoggedAt  time.Time `bson:"loggedAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
