package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schoolMaterial/internal/models"
)

func (h *Handler) CreateMaterial(c *gin.Context) {
	var material models.Material

	if err := c.BindJSON(&material); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uuid, err := h.service.CreateMaterial(material)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": uuid})

}

func (h *Handler) GetMaterialByUUID(c *gin.Context) {
	uuid := c.Param("uuid")

	material, err := h.service.GetMaterialByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, material)
}

func (h *Handler) UpdateMaterial(c *gin.Context) {
	uuid := c.Param("uuid")

	updateRequest := &models.UpdateRequest{}

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	material, err := h.service.UpdateMaterial(uuid, *updateRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, material)

}

func (h *Handler) GetAllMaterials(c *gin.Context) {
	active := c.Query("active") == "true"
	materialType := c.Query("type")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	materials, err := h.service.GetAllMaterials(active, materialType, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, materials)
}
