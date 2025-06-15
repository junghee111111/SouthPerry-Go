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
	"time"
)

func InsertAccount(email string, hashedPassword string) {
	newAccount := model.Account{
		Email:        email,
		PasswordHash: hashedPassword,
		LastLoggedAt: time.Now(),
		LastLoggedIp: "127.0.0.1",
		IsLoggedIn:   false,
		Sex:          false,
		Birthday:     time.Now(),
		IsBanned:     false,
	}
	result := db.MariaDB.Select("Email", "PasswordHash", "LastLoggedAt", "LastLoggedIp").Create(&newAccount)

	if result.Error != nil {
		log.Println("InsertAccount Error!", email, hashedPassword)
		log.Printf("  ==> %v\n", result.Error)
		return
	}

}

func FindAccountByEmail(email string) (model.Account, error) {
	account := &model.Account{}

	result := db.MariaDB.First(&account, "email = ?", email)
	if result.Error != nil {
		return *account, result.Error
	}

	return *account, nil
}
