package repository

import (
	"context"
	"errors"

	"github.com/tarkovskynik/Golang-ninja-project-2/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	NotFoundUser = errors.New("user doesn't exist by mail")
)

type UserInterface interface {
	Create(ctx context.Context, user *domain.User) error
	GetByEmailAndPassword(ctx context.Context, email, password string) (*domain.User, error)
	UpdateByEmail(ctx context.Context, email string, user *domain.User) error
	DeleteByEmail(ctx context.Context, email string) error
}

type User struct {
	con            *mongo.Database
	userCollection string
}

func NewUser(database *mongo.Database, userCollection string) *User {
	return &User{
		con:            database,
		userCollection: userCollection,
	}
}

func (u *User) Create(ctx context.Context, user *domain.User) error {
	_, err := u.con.Collection(u.userCollection).InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetByEmailAndPassword(ctx context.Context, email, password string) (*domain.User, error) {
	result := u.con.Collection(u.userCollection).FindOne(ctx, bson.M{"email": email, "password": password})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var user domain.User
	err := result.Decode(&user)

	return &user, err
}

func (u *User) UpdateByEmail(ctx context.Context, email string, user *domain.User) error {
	result, err := u.con.Collection(u.userCollection).UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return NotFoundUser
	}

	return nil
}

func (u *User) DeleteByEmail(ctx context.Context, email string) error {
	result, err := u.con.Collection(u.userCollection).DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return NotFoundUser
	}

	return nil
}
