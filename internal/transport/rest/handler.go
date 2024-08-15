package handler

import (
	"github.com/gin-gonic/gin"
	service "schoolMaterial/internal/services"
)

type Handler struct {
	service *service.MaterialService
}

func NewHandler(s *service.MaterialService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/materials", h.CreateMaterial)
	router.GET("/materials/:uuid", h.GetMaterialByUUID)

	return router

}
