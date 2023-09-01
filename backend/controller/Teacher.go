package controller

import (
	"github.com/B6332907/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Teacher

func CreateTeacher(c *gin.Context) {

	var Teacher entity.Teacher
	if err := c.ShouldBindJSON(&Teacher); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&Teacher).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teacher})

}

// GET /Teacher/:id

func GetTeacher(c *gin.Context) {

	var Teacher entity.Teacher

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Teachers WHERE id = ?", id).Scan(&Teacher).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teacher})

}

// GET /Teacher

func ListTeacher(c *gin.Context) {

	var Teacher []entity.Teacher

	if err := entity.DB().Raw("SELECT * FROM Teachers").Scan(&Teacher).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teacher})

}

// DELETE /Teacher/:id

func DeleteTeacher(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Teachers WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /Teacher

func UpdateTeacher(c *gin.Context) {

	var Teacher entity.Teacher

	if err := c.ShouldBindJSON(&Teacher); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", Teacher.ID).First(&Teacher); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	if err := entity.DB().Save(&Teacher).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teacher})

}
