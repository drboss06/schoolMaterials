package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schoolMaterial/internal/models"
	"schoolMaterial/pkg/logger"
)

func (h *Handler) CreateMaterial(c *gin.Context) {
	logger.GetLogger().Info("Create Material")
	var material models.Material

	if err := c.BindJSON(&material); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.GetLogger().Error(err)
		return
	}

	uuid, err := h.service.CreateMaterial(material)

	logger.GetLogger().Info("call CreateMaterial service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.GetLogger().Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": uuid})

}

func (h *Handler) GetMaterialByUUID(c *gin.Context) {
	logger.GetLogger().Info("Get Material By UUID")

	uuid := c.Param("uuid")

	material, err := h.service.GetMaterialByUUID(uuid)
	logger.GetLogger().Info("call GetMaterialByUUID service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.GetLogger().Error(err)
		return
	}

	c.JSON(http.StatusOK, material)
}

func (h *Handler) UpdateMaterial(c *gin.Context) {
	logger.GetLogger().Info("Update Material")

	uuid := c.Param("uuid")

	updateRequest := &models.UpdateRequest{}

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.GetLogger().Error(err)
		return
	}

	material, err := h.service.UpdateMaterial(uuid, *updateRequest)

	logger.GetLogger().Info("call UpdateMaterial service")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.GetLogger().Error(err)
		return
	}

	c.JSON(http.StatusOK, material)

}

func (h *Handler) GetAllMaterials(c *gin.Context) {
	logger.GetLogger().Info("Get All Materials")
	active := c.Query("active") == "true"
	materialType := c.Query("type")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	materials, err := h.service.GetAllMaterials(active, materialType, startDate, endDate)
	logger.GetLogger().Info("call GetAllMaterials service")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.GetLogger().Error(err)
		return
	}

	c.JSON(http.StatusOK, materials)
}
