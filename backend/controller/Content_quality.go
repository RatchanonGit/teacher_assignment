package controller

import (
	"github.com/B6332907/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Content_quality

func CreateContent_quality(c *gin.Context) {

	var Content_quality entity.Content_quality
	if err := c.ShouldBindJSON(&Content_quality); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&Content_quality).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Content_quality})

}

// GET /Content_quality/:id

func GetContent_quality(c *gin.Context) {

	var Content_quality entity.Content_quality

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Content_qualities WHERE id = ?", id).Scan(&Content_quality).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Content_quality})

}

// GET /Content_quality

func ListContent_quality(c *gin.Context) {

	var Content_quality []entity.Content_quality

	if err := entity.DB().Raw("SELECT * FROM Content_qualities").Scan(&Content_quality).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Content_quality})

}

// DELETE /Content_quality/:id

func DeleteContent_quality(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Content_qualitls WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /Teaching_duration

func UpdateContent_quality(c *gin.Context) {

	var Content_quality entity.Content_quality

	if err := c.ShouldBindJSON(&Content_quality); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", Content_quality.ID).First(&Content_quality); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	if err := entity.DB().Save(&Content_quality).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Content_quality})

}
