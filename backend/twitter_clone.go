package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	db *sqlx.DB
}

func main() {
	USERDB := os.Getenv("DB_USER")
	PASSDB := os.Getenv("DB_PASS")
	HOSTDB := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	os.Getenv("")
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", 
	HOSTDB, 5432, USERDB, PASSDB, DBNAME, "disable")
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	server := Server{db}
	router := setupRouter(&server)
	log.Println("Starting Server")
	router.Run(":5000")

}

func setupRouter(server *Server) *gin.Engine {
	router := gin.Default()
	router.GET("/tweet/:id", server.getTweet)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router
}

func (s *Server) getTweet(c *gin.Context) {
	id := c.Param("id")
	var tweet Tweet
	err := s.db.Get(&tweet, fmt.Sprintf("SELECT * FROM tweets WHERE id=%s", id))

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	c.JSON(http.StatusOK, tweet)
}

func postTweet(c *gin.Context) {

}

type Tweet struct {
	ID      int    `json:"id"`
	USER_ID string `json:"user_id"`
	Content string `json:"content"`
}

type User struct {
}
