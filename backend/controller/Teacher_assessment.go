package controller

import (
	"github.com/B6332907/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Teacher_assessments

func CreateTeacher_assessment(c *gin.Context) {

	var Teacher_assessment entity.Teacher_assessment
	if err := c.ShouldBindJSON(&Teacher_assessment); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&Teacher_assessment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teacher_assessment})

}

// GET /user/:id

func GetTeacher_assessment(c *gin.Context) {

	var Teacher_assessment entity.Teacher_assessment

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Teacher_assessments WHERE id = ?", id).Scan(&Teacher_assessment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teacher_assessment})

}

// GET /Teacher_assessments

func ListTeacher_assessment(c *gin.Context) {

	var Teacher_assessment []entity.Teacher_assessment

	if err := entity.DB().Preload("Student").Preload("Teacher").Preload("Teaching_duration").Preload("Content_difficulty_level").Preload("Content_quality").Raw("SELECT * FROM Teacher_assessments").Find(&Teacher_assessment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teacher_assessment})

}

// DELETE /Teacher_assessments/:id

func DeleteTeacher_assessment(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Teacher_assessments WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /Teacher_assessments

func UpdateTeacher_assessment(c *gin.Context) {

	var Teacher_assessment entity.Teacher_assessment

	if err := c.ShouldBindJSON(&Teacher_assessment); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", Teacher_assessment.ID).First(&Teacher_assessment); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	if err := entity.DB().Save(&Teacher_assessment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Teacher_assessment})

}
