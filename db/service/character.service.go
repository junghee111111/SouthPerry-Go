/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package service

import (
	"SouthPerry/db/model"
	"SouthPerry/db/repository"
)

func CreateCharacter(accId uint, c *model.Character) {

	if accId < 1 {
		return
	}

	c.AccountID = accId

	repository.InsertCharacter(c)
}
