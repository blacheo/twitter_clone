package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	defer conn.Close()

	router := gin.Default()
	router.GET("/tweet/:id", getTweet)

	router.Run("localhost:5000")

}

func getTweet(c *gin.Context) {
	pgxpool.New(context.Background())
	tweet_id := c.Param("id")

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
