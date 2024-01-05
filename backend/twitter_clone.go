package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db *sql.DB
}

func main() {
	UNAMEDB := os.Getenv("UNAMEDB")
	PASSDB := os.Getenv("PASSDB")
	HOSTDB := os.Getenv("HOSTDB")
	DBNAME := os.Getenv("DBNAME")
	os.Getenv("")
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", UNAMEDB, PASSDB, HOSTDB, DBNAME)
	db, err := sql.Open("postgres", connStr)
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
	id := c.Query("id")
	var tweet Tweet
	err := s.db.QueryRow(fmt.Sprintf("SELECT * FROM tweets WHERE id=%s", id)).Scan(&tweet)

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
