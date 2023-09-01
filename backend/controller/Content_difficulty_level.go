package controller

import (
	"github.com/B6332907/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Content_difficulty_level

func CreateContent_difficulty_level(c *gin.Context) {

	var Content_difficulty_level entity.Content_difficulty_level
	if err := c.ShouldBindJSON(&Content_difficulty_level); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&Content_difficulty_level).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Content_difficulty_level})

}

// GET /Content_difficulty_level/:id

func GetContent_difficulty_level(c *gin.Context) {

	var Content_difficulty_level entity.Content_difficulty_level

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Content_difficulty_levels WHERE id = ?", id).Scan(&Content_difficulty_level).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Content_difficulty_level})

}

// GET /Content_difficulty_level

func ListContent_difficulty_level(c *gin.Context) {

	var Content_difficulty_level []entity.Content_difficulty_level

	if err := entity.DB().Raw("SELECT * FROM Content_difficulty_levels").Scan(&Content_difficulty_level).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Content_difficulty_level})

}

// DELETE /Content_difficulty_level/:id

func DeleteContent_difficulty_level(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Content_difficulty_levels WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /Teaching_duration

func UpdateContent_difficulty_level(c *gin.Context) {

	var Content_difficulty_level entity.Content_difficulty_level

	if err := c.ShouldBindJSON(&Content_difficulty_level); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", Content_difficulty_level.ID).First(&Content_difficulty_level); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	if err := entity.DB().Save(&Content_difficulty_level).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Content_difficulty_level})

}
