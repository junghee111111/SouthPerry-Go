/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Character struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	accountId int                `bson:"accountId"`
	Name      string             `bson:"name"`
}
