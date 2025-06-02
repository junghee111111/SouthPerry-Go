/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package repository

import (
	"SouthPerry/db"
	"SouthPerry/db/model"
	"context"
	"log"
)

func InsertCharacter(ctx context.Context, c *model.Character) {
	cAccount := db.DB.Collection("characters")
	_, err := cAccount.InsertOne(ctx, &c)

	if err != nil {
		log.Println("InsertCharacter Error!")
		log.Printf("  ==> %v\n", err)
		return
	}

}
