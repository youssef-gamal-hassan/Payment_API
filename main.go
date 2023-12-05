package main

import (
	_ "errors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
)

type cardData struct {
	CardNumber     uint    `json:"cardNumber"`
	CVV            uint    `json:"cvv"`
	ExpirationDate [3]uint `json:"expirationDate"`
}

var cards = []cardData{
	{CardNumber: 123456789123, CVV: 123, ExpirationDate: [3]uint{11, 1, 24}},
}

func getCards(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cards)
}

func main() {
	router := gin.Default()
	router.GET("/cards", getCards)
	router.Run("localhost:8080")
}
