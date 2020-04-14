package main


import (
	"net/http"

	"github.com/gin-gonic/gin"

	middleware "github.com/s12i/gin-throttle"

	cors "github.com/gin-contrib/cors"

	"github.com/jeffdhooton/dadjokes.lol/controllers"

	"github.com/jeffdhooton/dadjokes.lol/models"
)

func main() {
	router := gin.Default()

	db := models.SetupModels()

	router.Use(middleware.Throttle(2, 2))

	router.LoadHTMLGlob("templates/*")

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
	}))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Test Title",
		})
	})

	// router.GET("/jokes", controllers.FindJokes)
	router.GET("/random", controllers.RandomJoke)
	// router.POST("/jokes", controllers.CreateJoke)
	// router.GET("/jokes/:id", controllers.FindJoke)
	// router.PATCH("/jokes/:id", controllers.UpdateJoke)
	// router.DELETE("jokes/:id", controllers.DeleteJoke)

	router.Run()
}
