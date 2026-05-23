package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Project struct {
	Name        string
	Description string
}

var projects = []Project{
	{
		Name:        "Personal Website",
		Description: "Website to host personal content and projects.",
	},
	{
		Name:        "Music Generating Neural Network",
		Description: "Generative Adversarial Network and Recurrent Neural Network that produced music after learning on MIDI data.",
	},
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"Title":    "Home",
			"Projects": projects,
		})
	})

	router.Run(":8080")
}
