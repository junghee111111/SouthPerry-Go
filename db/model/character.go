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
	CharId    int                `bson:"charId"`
	AccountId int                `bson:"accountId"`
	Name      string             `bson:"name"`

	Face   int `bson:"face"`
	Hair   int `bson:"hair"`
	Top    int `bson:"top"`
	Bottom int `bson:"bottom"`
	Shoes  int `bson:"shoes"`
	Weapon int `bson:"weapon"`

	Str uint16 `bson:"str"`
	Dex uint16 `bson:"dex"`
	Int uint16 `bson:"int"`
	Luk uint16 `bson:"luk"`

	Level int `bson:"level" default:"1"`
	Job   int `bson:"job" default:"0"`
}
