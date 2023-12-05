package main

import (
	"encoding/json"
	_ "errors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
	"os"
	"path/filepath"
)

type cardData struct {
	CardNumber     uint    `json:"cardNumber"`
	CVV            uint    `json:"cvv"`
	ExpirationDate [3]uint `json:"expirationDate"`
}

type Cards struct {
	Cards []cardData `json:"cards"`
}

var cards = Cards{
	Cards: []cardData{
		{CardNumber: 123456789123, CVV: 123, ExpirationDate: [3]uint{11, 1, 24}},
	},
}

func getCards(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cards)
}

func addCard(c *gin.Context) {
	var newCard cardData

	if err := c.BindJSON(&newCard); err != nil {
		return
	}

	cards.Cards = append(cards.Cards, newCard)
	c.IndentedJSON(http.StatusCreated, newCard)
	cards.saveCards()
}

func (c *Cards) loadCards() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(filepath.Join(home, "Documents", "Y3T1", "SE", "Project", "API", "cards.json"))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
}

func (c *Cards) saveCards() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(home, "Documents", "Y3T1", "SE", "Project", "API", "cards.json"), data, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	cards.loadCards()
	router := gin.Default()
	router.GET("/cards", getCards)
	router.POST("/cards", addCard)
	router.Run("localhost:8080")

}
