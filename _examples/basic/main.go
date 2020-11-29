package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	chaos "github.com/msfidelis/gin-chaos-monkey"
)

type Healthcheck struct {
	Status int `json:"status" binding:"required"`
}

func main() {

	// Use this on your environment. This is a example only
	os.Setenv("CHAOS_MONKEY_ENABLED", "true")
	os.Setenv("CHAOS_MONKEY_LATENCY", "true")
	os.Setenv("CHAOS_MONKEY_LATENCY_MAX_TIME", "5000")
	os.Setenv("CHAOS_MONKEY_LATENCY_MIN_TIME", "1000")
	os.Setenv("CHAOS_MONKEY_MODE", "critical")

	router := gin.Default()

	//Middlewares
	router.Use(gin.Recovery())
	router.Use(chaos.Load())

	router.GET("/healthcheck", HealthcheckGet)

	router.Run()
}

func HealthcheckGet(c *gin.Context) {

	var response Healthcheck
	response.Status = http.StatusOK

	c.JSON(http.StatusOK, response)

}
