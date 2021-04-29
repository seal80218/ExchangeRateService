package main

import (
	"fmt"
	"net/http"

	"github.com/exange_rate_service/repository"
	schedule "github.com/exange_rate_service/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server start")

	go schedule.Run()

	router := gin.Default()

	router.GET("/exchangerate/:currency", func(c *gin.Context) {
		cur := c.Param("currency")

		repo := new(repository.ExchangeRateRepository)
		repo.Init()
		curData := repo.Get(cur)

		c.JSON(http.StatusOK, curData)
	})

	router.Run(":52001")
}
