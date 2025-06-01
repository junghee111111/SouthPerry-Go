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
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func InsertAccount(ctx context.Context, email string, hashedPassword string) {
	cAccount := db.DB.Collection("accounts")
	newAccount := model.Account{
		Email:        email,
		PasswordHash: hashedPassword,
	}
	_, err := cAccount.InsertOne(ctx, newAccount)

	if err != nil {
		log.Println("InsertAccount Error!", email, hashedPassword)
		log.Printf("  ==> %v\n", err)
		return
	}

}

func FindAccountByEmail(ctx context.Context, email string) (model.Account, error) {
	cAccount := db.DB.Collection("accounts")
	var storedAccount model.Account
	err := cAccount.FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&storedAccount)

	if err != nil {
		log.Println("FindAccountByEmail Error!", email)
		log.Printf("  ==> %v\n", err)
		return model.Account{}, err
	}

	return storedAccount, nil
}

func IsNameUsed(ctx context.Context, name string) bool {
	foundCharacter := db.DB.Collection("characters")
	var storeCharacter model.Character
	err := foundCharacter.FindOne(ctx, bson.M{
		"name": name,
	}).Decode(&storeCharacter)
	return err == nil
}
