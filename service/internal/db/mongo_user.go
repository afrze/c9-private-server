package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type WebUser struct {
	AccNo        int64     `bson:"cAccNo"`
	AccId        string    `bson:"cAccId"`
	PasswordHash string    `bson:"cPasswordHash"`
	Created      time.Time `bson:"created"`
}

func InsertWebUser(ctx context.Context, mdb *mongo.Database, u WebUser) error {
	_, err := mdb.Collection("users").InsertOne(ctx, u)
	return err
}

func FindWebUser(ctx context.Context, mdb *mongo.Database, accId string) (WebUser, error) {
	var u WebUser
	err := mdb.Collection("users").FindOne(ctx, map[string]any{"cAccId": accId}).Decode(&u)
	return u, err
}
