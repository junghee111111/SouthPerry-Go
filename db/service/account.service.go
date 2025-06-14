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
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, _ := hashPassword(password)
	repository.InsertAccount(email, hashedPassword)
}

func CheckAccount(email string, password string) (result enum.AccountRespCode, account model.Account) {
	account, err := repository.FindAccountByEmail(email)

	if err != nil {
		return enum.CheckAccountResp.WrongID, model.Account{}
	}

	if checkPasswordHash(password, account.PasswordHash) {
		return enum.CheckAccountResp.Success, account
	} else {
		return enum.CheckAccountResp.WrongPassword, model.Account{}
	}
}

func CheckCharacterName(name string) bool {
	return repository.IsNameUsed(name)
}
