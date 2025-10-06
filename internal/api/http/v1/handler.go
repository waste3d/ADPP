package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waste3d/ADPP/internal/domain"
)

type JobStorageInterface interface {
	CreateJob(input int) (*domain.Job, error)
}

type Handler struct {
	storage JobStorageInterface
}

func NewHandler(storage JobStorageInterface) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) CreateJobHandler(c *gin.Context) {
	var requestBody CreateJobRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job, err := h.storage.CreateJob(requestBody.Input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := CreateJobResponse{
		ID: job.ID,
	}

	c.JSON(http.StatusCreated, response)
}
