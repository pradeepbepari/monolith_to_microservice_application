package handlers

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Createuser(ctx *gin.Context)
}
