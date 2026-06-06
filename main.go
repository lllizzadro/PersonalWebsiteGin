package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

var db *sql.DB = nil

type Project struct {
	Name        string
	Description string
}

type Message struct {
	Id        int
	Name      string
	Message   string
	Timestamp time.Time
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
	initDB()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"Title":    "Home",
			"Projects": projects,
		})
	})

	router.GET("/guestbook", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name, message, timestamp FROM guestbook ORDER BY timestamp DESC")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var messages []Message
		for rows.Next() {
			var m Message
			rows.Scan(&m.Id, &m.Name, &m.Message, &m.Timestamp)
			messages = append(messages, m)
		}
		c.HTML(http.StatusOK, "guestbook", gin.H{
			"Title":    "Guestbook",
			"Messages": messages,
		})
	})

	router.GET("/dice", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dice", gin.H{
			"Title": "Dice",
		})
	})

	router.POST("/guestbook", func(c *gin.Context) {
		name := c.PostForm("name")
		message := c.PostForm("message")
		_, err := db.Exec("INSERT INTO guestbook (name, message) VALUES (?, ?)", name, message)
		if err != nil {
			log.Fatal(err)
		}
		c.Redirect(http.StatusSeeOther, "/guestbook")
	})

	router.Run(":8080")
}

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "guestbook.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS guestbook (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, message TEXT NOT NULL, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	}
	return
}
