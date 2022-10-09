package service

import (
	"context"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/tarkovskynik/Golang-ninja-project-2/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileRepository interface {
	Create(ctx context.Context, file domain.File) error
	List(ctx context.Context, userID primitive.ObjectID) ([]domain.File, error)
}

type File struct {
	repository  FileRepository
	minioClient *minio.Client
}

func NewFile(repo FileRepository, mc *minio.Client) *File {
	return &File{
		repository:  repo,
		minioClient: mc,
	}
}

func (f *File) UploadFile(ctx context.Context, bucketName string, input domain.UploadInput) error {
	if err := f.NewBucket(ctx, bucketName); err != nil {
		return err
	}

	_, err := f.minioClient.PutObject(
		ctx,
		input.Bucket,
		input.Name,
		input.Payload,
		input.Size,
		minio.PutObjectOptions{ContentType: "image/png"},
	)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) GenerateFileURL(ctx context.Context, bucketname, filename string) (string, error) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+filename+"\"")

	presignedURL, err := f.minioClient.PresignedGetObject(ctx, bucketname, filename, time.Hour*24, reqParams)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}

func (f *File) NewBucket(ctx context.Context, bucketName string) error {

	found, err := f.minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	if !found {
		f.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}

	return nil
}

func (f *File) Create(ctx context.Context, file domain.File) error {
	return f.repository.Create(ctx, file)
}

func (f *File) List(ctx context.Context, userID primitive.ObjectID) ([]domain.File, error) {
	return f.repository.List(ctx, userID)
}
