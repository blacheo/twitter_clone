package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	db *pgxpool.Pool
}

func main() {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	defer conn.Close()

	server := Server{conn}
	router := gin.Default()
	router.GET("/tweet/:id", server.getTweet)

	router.Run("localhost:5000")

}

func (s *Server) getTweet(c *gin.Context) {
	id := c.Query("id")
	var tweet Tweet
	err := s.db.QueryRow(context.Background(), fmt.Sprintf("SELECT * FROM tweets WHERE id=%s", id)).Scan(&tweet)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	c.JSON(http.StatusOK, tweet)
}

func postTweet(c *gin.Context) {

}

type Tweet struct {
	ID      int    `json:"id"`
	USER_ID int    `json:"user_id"`
	Content string `json:"content"`
}

type User struct {
}
