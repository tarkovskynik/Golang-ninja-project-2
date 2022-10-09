package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tarkovskynik/Golang-ninja-project-2/internal/domain"
	"github.com/tarkovskynik/Golang-ninja-project-2/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) getFiles(c *gin.Context) {
	userIDRow := c.Request.Header.Get("user_id")
	if userIDRow == "" {
		logger.LogrusLogger(domain.NotFoundUserID, "getFiles", domain.NotFoundUserID.Error())
		NewResponse(c, http.StatusBadRequest, domain.NotFoundUserID.Error())
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIDRow)
	if err != nil {
		logger.LogrusLogger(err, "getFiles", err.Error())
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	files, err := h.file.List(c.Request.Context(), userID)
	if err != nil {
		logger.LogrusLogger(err, "getFiles", err.Error())
		NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, files)
}

// TODO content type check
func (h *Handler) uploadFile(c *gin.Context) {
	userID := c.Request.Header.Get("user_id")
	if userID == "" {
		logger.LogrusLogger(domain.NotFoundUserID, "getFiles", domain.NotFoundUserID.Error())
		NewResponse(c, http.StatusBadRequest, domain.NotFoundUserID.Error())
		return
	}

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10*1024*1024)
	src, hdr, err := c.Request.FormFile("file")
	if err != nil {
		logger.LogrusLogger(err, "uploadFile", err.Error())
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer src.Close()

	file := domain.UploadInput{
		Name:    hdr.Filename,
		Bucket:  userID,
		Payload: src,
		Size:    hdr.Size,
	}

	if err = h.file.UploadFile(c.Request.Context(), userID, file); err != nil {
		logger.LogrusLogger(err, "uploadFile", err.Error())
		NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	url, err := h.file.GenerateFileURL(c.Request.Context(), file.Bucket, file.Name)
	if err != nil {
		logger.LogrusLogger(err, "uploadFile", err.Error())
		NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		logger.LogrusLogger(err, "uploadFile", err.Error())
		NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fileInfo := domain.File{
		UserID:   objID,
		URL:      url,
		Size:     uint64(file.Size),
		LoadedAt: time.Now(),
	}

	if err = h.file.Create(c.Request.Context(), fileInfo); err != nil {
		logger.LogrusLogger(err, "uploadFile", err.Error())
		NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}
