package rest

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tarkovskynik/Golang-ninja-project-2/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File interface {
	UploadFile(ctx context.Context, bucketName string, input domain.UploadInput) error
	GenerateFileURL(ctx context.Context, bucketname, filename string) (string, error)
	Create(ctx context.Context, file domain.File) error
	List(ctx context.Context, userID primitive.ObjectID) ([]domain.File, error)
	NewBucket(ctx context.Context, bucketName string) error
}

type Handler struct {
	file File
}

func NewHandler(file File) *Handler {
	return &Handler{
		file: file,
	}
}

type statusResponse struct {
	Message string `json:"status"`
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.Default()

	filesApi := router.Group("/store")
	{
		filesApi.PUT("/upload", h.uploadFile)
		filesApi.GET("/files", h.getFiles)
	}

	return router
}

