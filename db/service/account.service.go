/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package service

import (
	"SouthPerry/db/enum"
	"SouthPerry/db/model"
	"SouthPerry/db/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func CheckAccount(email string, password string) (result enum.AccountRespCode, account model.Account) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	account, err := repository.FindAccountByEmail(ctx, email)

	if err != nil {
		return enum.CheckAccountResp.WrongID, model.Account{}
	}

	if checkPasswordHash(password, account.PasswordHash) {
		return enum.CheckAccountResp.Success, account
	} else {
		return enum.CheckAccountResp.WrongPassword, model.Account{}
	}
}
