package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	chaos "github.com/msfidelis/gin-chaos-monkey"
)

type Healthcheck struct {
	Status      int    `json:"status" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func main() {
	router := gin.Default()

	//Enable Chaos Monkey in Specific Route
	router.GET("/healthcheck/chaos", chaos.Load(), HealthcheckGet)

	router.GET("/healthcheck", HealthcheckGet)

	router.Run()
}

func HealthcheckGet(c *gin.Context) {

	var response Healthcheck
	response.Status = http.StatusOK
	response.Description = "default"

	c.JSON(http.StatusOK, response)

}
