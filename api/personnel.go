package api

import (
	db "boba/db/sqlc"
	"boba/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type createPersonnelRequest struct {
	Name string `json:"name" binding:"required,alphanum"`
	Shop int64  `json:"shop" binding:"required"`
	Pin  string `json:"pin" binding:"required,number"`
}

type createPersonnelResponse struct {
	BadgeNumber int64       `json:"badge_number"`
	Name        string      `json:"name"`
	Shop        int64       `json:"shop"`
	Email       pgtype.Text `json:"email"`
	CreatedAt   time.Time   `json:"created_at"`
}

func (server *Server) createPersonnel(ctx *gin.Context) {
	var req createPersonnelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPin, err := util.HashPin(req.Pin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreatePersonnelParams{
		Name:      req.Name,
		Shop:      req.Shop,
		HashedPin: hashedPin,
	}

	personnel, err := server.store.CreatePersonnel(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createPersonnelResponse{
		BadgeNumber: personnel.BadgeNumber,
		Name:        personnel.Name,
		Shop:        personnel.Shop,
		Email:       personnel.Email,
		CreatedAt:   personnel.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)

}
