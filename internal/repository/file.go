package repository

import (
	"context"

	"github.com/tarkovskynik/Golang-ninja-project-2/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type File struct {
	con            *mongo.Database
	fileCollection string
}

func NewFile(database *mongo.Database, fileCollection string) *File {
	return &File{
		con:            database,
		fileCollection: fileCollection,
	}
}

func (f *File) Create(ctx context.Context, file domain.File) error {
	_, err := f.con.Collection(f.fileCollection).InsertOne(ctx, file)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) List(ctx context.Context, userID primitive.ObjectID) ([]domain.File, error) {
	filter := bson.D{{Key: "user_id", Value: userID}}

	cur, err := f.con.Collection(f.fileCollection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var filesList []domain.File
	for cur.Next(ctx) {
		var file domain.File
		if err := cur.Decode(&file); err != nil {
			return nil, err
		}
		filesList = append(filesList, file)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return filesList, nil
}
