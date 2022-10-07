package domain

import (
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID       primitive.ObjectID `json:"id"        bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"user_id"   bson:"user_id"`
	URL      string             `json:"url"       bson:"url"`
	Size     uint64             `json:"size"      bson:"size"`
	LoadedAt time.Time          `json:"loaded_at" bson:"loaded_at"`
}

type UploadInput struct {
	Name    string
	Bucket  string
	Payload multipart.File
	Size    int64
}
