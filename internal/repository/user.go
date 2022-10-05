package repository

import (
	"context"

	"github.com/tarkovskynik/Golang-ninja-project-2/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	filter := bson.M{"email": email, "password": password}
	var user domain.User
	err := u.con.Collection(u.userCollection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (u *User) UpdateByEmail(ctx context.Context, email string, user *domain.User) error {
	filter := bson.M{"email": email}
	update := bson.M{"$set": user}
	queryResult, err := u.con.Collection(u.userCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if queryResult.ModifiedCount == 0 {
		return domain.NotFoundUser
	}

	return nil
}

func (u *User) DeleteByEmail(ctx context.Context, email string) error {
	queryResult, err := u.con.Collection(u.userCollection).DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		return err
	}
	if queryResult.DeletedCount == 0 {
		return domain.NotFoundUser
	}

	return nil
}
