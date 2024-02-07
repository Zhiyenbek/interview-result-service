package handler

import (
	"github.com/Zhiyenbek/interview-result-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *handler) CreateResult(c *gin.Context) {
	req := &models.CreateResultRequest{}
	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		h.logger.Errorf("ERROR: invalid input, some fields are incorrect: %s\n", err.Error())
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
		return
	}

	err := h.service.ResultService.CreateResult(req)
	if err != nil {
		h.logger.Errorf("Error occurred while login: %v", err)
		c.JSON(500, sendResponse(-1, nil, models.ErrInternalServer))
		return
	}
	c.JSON(201, sendResponse(0, nil, nil))
}
