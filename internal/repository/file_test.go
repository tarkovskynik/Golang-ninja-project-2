package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tarkovskynik/Golang-ninja-project-2/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type FileInterface interface {
	Create(ctx context.Context, file domain.File) error
	List(ctx context.Context, userID primitive.ObjectID) ([]domain.File, error)
}

func createFile(t *testing.T, file domain.File) {
	err := TestFileRepository.Create(Ctx, file)
	require.NoError(t, err)
}

func getList(t *testing.T, filesAct []domain.File, id primitive.ObjectID) {
	files, err := TestFileRepository.List(Ctx, id)
	require.NoError(t, err)
	require.Equal(t, filesAct, files)
}

func TestFile_Create(t *testing.T) {
	id := primitive.NewObjectID()
	user_id := primitive.NewObjectID()
	file := domain.File{
		ID:       id,
		UserID:   user_id,
		URL:      "test@test.com",
		Size:     1024,
		LoadedAt: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC),
	}

	createFile(t, file)
}

func TestFile_List(t *testing.T) {
	id1, id2 := primitive.NewObjectID(), primitive.NewObjectID()
	user_id := primitive.NewObjectID()

	files := []domain.File{
		{
			ID:       id1,
			UserID:   user_id,
			URL:      "testList1",
			Size:     1024,
			LoadedAt: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			ID:       id2,
			UserID:   user_id,
			URL:      "testList2",
			Size:     1024,
			LoadedAt: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC),
		},
	}

	for _, file := range(files) {
		createFile(t, file)
	}

	getList(t, files, user_id)
}
