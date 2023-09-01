package main

import (
	"github.com/B6332907/sa-65-example/entity"

	"github.com/B6332907/sa-65-example/controller"

	"github.com/B6332907/sa-65-example/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/") 
	
{
	router.Use(middlewares.Authorizes())
	{
		// Teacher_assessment
		r.GET("/Teacher_assessments", controller.ListTeacher_assessment)
		r.GET("/Teacher_assessment/:id", controller.GetTeacher_assessment)
		r.POST("/Teacher_assessments", controller.CreateTeacher_assessment)
		r.PATCH("/Teacher_assessments", controller.UpdateTeacher_assessment)
		r.DELETE("/Teacher_assessments/:id", controller.DeleteTeacher_assessment)

		//Student
		r.GET("/Students", controller.ListStudent)
		r.GET("/Student/:id", controller.GetStudent)
		r.POST("/Student", controller.CreateStudent)
		r.PATCH("/Student", controller.UpdateStudent)
		r.DELETE("/Student/:id", controller.DeleteStudent)

		//Teacher
		r.GET("/teachers", controller.ListTeacher)
		r.GET("/teacher/:id", controller.GetTeacher)
		r.POST("/Teacher", controller.CreateTeacher)
		r.PATCH("/Teacher", controller.UpdateTeacher)
		r.DELETE("/Teacher/:id", controller.DeleteTeacher)

		//Teaching_duration
		r.GET("/Teaching_durations", controller.ListTeaching_duration)
		r.GET("/Teaching_duration/:id", controller.GetTeaching_duration)
		r.POST("/Teaching_duration", controller.CreateTeaching_duration)
		r.PATCH("/Teaching_duration", controller.UpdateTeaching_duration)
		r.DELETE("/Teaching_duration/:id", controller.DeleteTeaching_duration)

		//Content_difficulty_level
		r.GET("/Content_difficulty_levels", controller.ListContent_difficulty_level)
		r.GET("/Content_difficulty_level/:id", controller.GetContent_difficulty_level)
		r.POST("/Content_difficulty_level", controller.CreateContent_difficulty_level)
		r.PATCH("/Content_difficulty_level", controller.UpdateContent_difficulty_level)
		r.DELETE("/Content_difficulty_level/:id", controller.DeleteContent_difficulty_level)

		//Content_quality
		r.GET("/Content_qualitys", controller.ListContent_quality)
		r.GET("/Content_quality/:id", controller.GetContent_quality)
		r.POST("/Content_quality", controller.CreateContent_quality)
		r.PATCH("/Content_quality", controller.UpdateContent_quality)
		r.DELETE("/Content_quality/:id", controller.DeleteContent_quality)
	// Run the server
		}

	}
	// Signup User Route
	r.POST("/signup", controller.CreateStudent)
	// login User Route
	r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run()

}



func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return
		}

		c.Next()
	}

}
