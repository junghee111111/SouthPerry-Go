/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package repository

import (
	"SouthPerry/db"
	"SouthPerry/db/model"
	"log"
)

func InsertCharacter(c *model.Character) {
	newCharacter := c
	result := db.MariaDB.Create(&newCharacter)

	if result.Error != nil {
		log.Println("InsertCharacter Error!")
		log.Printf("  ==> %v\n", result.Error)
		return
	}

}

func IsNameUsed(name string) bool {
	result := db.MariaDB.Where("name = ?", name).First(&model.Character{})

	if result.Error != nil {
		log.Println("IsNameUsed Error!", name)
		log.Printf("  ==> %v\n", result.Error)
		return false
	}

	return true
}
