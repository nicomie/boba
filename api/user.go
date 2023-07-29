package api

import (
	db "boba/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name:    req.Name,
		Balance: 0,
	}

	acc, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, acc)

}
