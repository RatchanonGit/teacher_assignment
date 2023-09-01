package controller

import (
	"github.com/B6332907/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Teaching_duration

func CreateTeaching_duration(c *gin.Context) {

	var Teaching_duration entity.Teaching_duration
	if err := c.ShouldBindJSON(&Teaching_duration); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&Teaching_duration).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teaching_duration})

}

// GET /Teaching_duration/:id

func GetTeaching_duration(c *gin.Context) {

	var Teaching_duration entity.Teaching_duration

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Teaching_durations WHERE id = ?", id).Scan(&Teaching_duration).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teaching_duration})

}

// GET /Teaching_duration

func ListTeaching_duration(c *gin.Context) {

	var Teaching_duration []entity.Teaching_duration

	if err := entity.DB().Raw("SELECT * FROM Teaching_durations").Scan(&Teaching_duration).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teaching_duration})

}

// DELETE /Teaching_duration/:id

func DeleteTeaching_duration(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Teaching_durations WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /Teaching_duration

func UpdateTeaching_duration(c *gin.Context) {

	var Teaching_duration entity.Teaching_duration

	if err := c.ShouldBindJSON(&Teaching_duration); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", Teaching_duration.ID).First(&Teaching_duration); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	if err := entity.DB().Save(&Teaching_duration).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teaching_duration})

}
