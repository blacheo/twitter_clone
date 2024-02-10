package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	router.GET("/tweets", server.getTopTweets)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Authorization Group
	authorized := router.Group("/")
	authorized.Use(AuthMiddleware())
	{
		authorized.POST("/like/:tweet_id", server.likeTweet)
	}

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

func (s *Server) getTopTweets(c *gin.Context) {
	const nb_tweets = "50"
	tweets := []Tweet{}
	err := s.db.Select(&tweets, "SELECT * FROM tweets ORDER BY id ASC LIMIT $1", nb_tweets)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	c.JSON(http.StatusOK, tweets)
}
func (s *Server) postTweet(c *gin.Context) {

}

func (s *Server) likeTweet(c *gin.Context) {
	tweet_id := c.Query("tweet_id")
	user_id := c.Query("user_id")
	_, err := s.db.Exec("INSERT INTO tweets (id, user_id, content) VALUES ($1, $2)", tweet_id, user_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LikeTweet failed: %v\n", err)
		c.Status(http.StatusNoContent)
		return
	}
	c.Status(http.StatusOK)
}

type Tweet struct {
	ID      int    `json:"id"`
	USER_ID string `json:"user_id"`
	Content string `json:"content"`
}

type User struct {
}
