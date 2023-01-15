package Player

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message" : "status ok",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"name" : "Lionel Messi",
		"nationality" : "Indonesia",
	})
}

func IdPlayerHandler(c *gin.Context){
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id" : id,
		"title" : title,
	})
}

func QueryHandler(c *gin.Context) {
	query := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title" : query,
		"price" : price,
	})
}

func CreatePlayer(c *gin.Context){
	var players Players
	
	err := c.ShouldBindJSON(&players)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"Name" : players.Name,
		"Nationality" : players.Nationality,
		"Age" : players.Age,
	})
}