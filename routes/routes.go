package routes

import (
	"net/http"
	"pov_golang/handlers"
	"pov_golang/middlerware"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	UserHandler handlers.UserHandler
}

func ApiRoutes(di Dependencies, r *gin.Engine) {
	router := r.Group("/api")
	router.GET("/health", HealthCheck)
	router.POST("/user", di.UserHandler.Createuser)
	router.Use(middlerware.Authenticate)
}

// @Health check endpoint
// @Summary Health Check
// @Description Get the health status of the API
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
