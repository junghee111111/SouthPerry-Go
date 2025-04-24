/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package service

import (
	"SouthPerry/db/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	BAN_ACCOUNT         = 3
	WRONG_PASSWORD      = 4
	WRONG_ID            = 5
	SYSTEM_ERROR        = 6
	ALREADY_LOGGEDIN    = 7
	SERVICE_UNAVAILABLE = 10
	OLDER20             = 11
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateAccount(email string, password string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hashedPassword, _ := hashPassword(password)
	repository.InsertAccount(ctx, email, hashedPassword)
}

func CheckAccount(email string, password string) (result uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	account, err := repository.FindAccountByEmail(ctx, email)

	if err != nil {
		return WRONG_ID
	}

	if checkPasswordHash(password, account.PasswordHash) {
		return OLDER20
	} else {
		return WRONG_PASSWORD
	}
}
