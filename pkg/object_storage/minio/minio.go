package minio

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ConnectionInfo struct {
	Url        string
	User       string
	Password   string
	SSLMode    bool
	BucketName string
}

func NewClient(info *ConnectionInfo) (*minio.Client, error) {
	client, err := minio.New(info.Url, &minio.Options{
		Creds:  credentials.NewStaticV4(info.User, info.Password, ""),
		Secure: info.SSLMode,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return client, err
}
