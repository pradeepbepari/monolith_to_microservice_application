package handlers

import (
	"net/http"
	"pov_golang/models"
	"pov_golang/service"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type handler struct {
	service service.UserService
	nrApp   newrelic.Application
}

func NewHandler(service service.UserService, nrApp newrelic.Application) UserHandler {
	return handler{
		service: service,
		nrApp:   nrApp,
	}
}

// CreateUser handles user signup
// @Summary Create a new user
// @Description Registers a new user and returns a JWT token
// @Tags Public
// @Accept json
// @Produce json
// @Param user body models.Users true "User Data"
// @Success 200
// @Failure 400 "{\"error\": \"invalid request body\"}"
// @Failure 500 "{\"error\": \"something went wrong\"}"
// @Router /user [post]
func (h handler) Createuser(ctx *gin.Context) {
	tx := h.nrApp.StartTransaction("Createuser")
	defer tx.End()
	var request models.Users
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	c := newrelic.NewContext(ctx.Request.Context(), tx)
	resp, err := h.service.Create(c, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": resp})
}
