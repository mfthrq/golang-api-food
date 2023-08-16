package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type fruit struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	Flavor   string `json:"flavor"`
	Quantity int    `json:"quantity"`
}

var fruits = []fruit{
	{
		ID:       "1",
		Name:     "Durian",
		Color:    "Green",
		Flavor:   "Sweet",
		Quantity: 1,
	},
	{
		ID:       "2",
		Name:     "Apple",
		Color:    "Red",
		Flavor:   "Sour",
		Quantity: 3,
	},
	{
		ID:       "3",
		Name:     "Banana",
		Color:    "Yellow",
		Flavor:   "Sweet",
		Quantity: 2,
	},
}

func getFruits(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fruits)
}

// func addFruits(c *gin.Context){
// 	var newFruit fruit

// 	if err := nil 
// }

func main() {
	router := gin.Default()
	router.GET("/fruits", getFruits)
	router.Run("localhost:8080")
}
