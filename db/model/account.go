/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	Characters   []Character `gorm:"constraint:OnDelete:CASCADE;"`
	Email        string      `gorm:"unique;uniqueIndex:idx_email"`
	PasswordHash string
	Sex          bool
	Birthday     time.Time
	IsBanned     bool
	IsLoggedIn   bool
	LastLoggedAt time.Time `gorm:"autoCreateTime"`
	LastLoggedIp string
}
