package domain

import "time"

type File struct {
	ID       int       `json:"id"        bson:"id"`
	UserID   int       `json:"user_id"   bson:"user_id"`
	URL      string    `json:"url"       bson:"url"`
	Size     int       `json:"size"      bson:"size"`
	LoadedAt time.Time `json:"loaded_at" bson:"loaded_at"`
}
