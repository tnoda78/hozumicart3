package main

import (
	"fmt"
	"image/gif"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tnoda78/newhozumicart3/generator"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.StaticFile("/", "views/index.html")
	router.GET("/image", func(c *gin.Context) {
		generator, err := generator.NewGenerator()
		if err != nil {
			c.String(http.StatusBadRequest, "ERROR.")
			return
		}

		color := c.DefaultQuery("color", "90CF4C")
		letter := c.DefaultQuery("letter", "")

		img, err := generator.GenerateImage(fmt.Sprintf("#%s", color), letter)
		if err != nil {
			c.String(http.StatusBadRequest, "ERROR.")
			return
		}

		gif.EncodeAll(c.Writer, img)
		c.Header("Content-Type", "image/gif")
	})

	router.Run(":" + port)
}
