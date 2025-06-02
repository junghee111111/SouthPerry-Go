/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package service

import (
	"SouthPerry/db/model"
	"SouthPerry/db/repository"
	"context"
	"time"
)

func CreateCharacter(accId int, c *model.Character) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if accId < 1 {
		return
	}

	c.AccountId = accId

	repository.InsertCharacter(ctx, c)
}
