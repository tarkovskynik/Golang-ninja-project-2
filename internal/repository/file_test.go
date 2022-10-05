package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tarkovskynik/Golang-ninja-project-2/internal/domain"
)

type FileInterface interface {
	Create(ctx context.Context, file domain.File) error
	List(ctx context.Context, userID int) ([]domain.File, error)
}

func createFile(t *testing.T, file domain.File) {
	err := TestFileRepository.Create(Ctx, file)
	require.NoError(t, err)
}

func getList(t *testing.T, filesAct []domain.File, id int) {
	files, err := TestFileRepository.List(Ctx, id)
	require.NoError(t, err)
	require.Equal(t, filesAct, files)
}

func TestFile_Create(t *testing.T) {
	file := domain.File{
		ID:       10002,
		UserID:   99,
		URL:      "test2",
		Size:     1024,
		LoadedAt: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC),
	}

	createFile(t, file)
}

func TestFile_List(t *testing.T) {
	id := 99
	files := []domain.File{
		{
			ID:       10001,
			UserID:   99,
			URL:      "test1",
			Size:     1024,
			LoadedAt: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			ID:       10002,
			UserID:   99,
			URL:      "test2",
			Size:     1024,
			LoadedAt: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC),
		},
	}
	getList(t, files, id)
}
