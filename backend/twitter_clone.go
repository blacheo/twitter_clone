package main

import (
	"net/http"
	"github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

func main() {


	router := gin.Default()
	router.GET("/tweets", getTweets)

	router.Run("localhost:5000")

}

var tweets = []Tweet{
	{ID: 1, Content: "Hello Twitter Clone!"},
	{ID: 2, Content: "Hi Everyone!"},
}

func getTweets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tweets)
}

type Tweet struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
