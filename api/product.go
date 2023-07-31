package api

import (
	db "boba/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createProductRequest struct {
	Name  string `json:"name" binding:"required,max=50"`
	Price int32  `json:"price" binding:"required,min=1"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse((err)))
		return
	}

	arg := db.CreateProductParams{
		Name:  req.Name,
		Price: req.Price,
	}

	product, err := server.store.CreateProduct(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse((err)))
		return
	}

	ctx.JSON(http.StatusOK, product)

}
