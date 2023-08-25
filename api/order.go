package api

import (
	db "boba/db/sqlc"
	"boba/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createOrderRequest struct {
	OrderItems []*OrderItem `json:"orders" binding:"dive"`
}

type OrderItem struct {
	ProductID int64 `json:"product_id" binding:"required,min=1"`
	Quantity  int32 `json:"quantity" binding:"required,min=1"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse((err)))
		return
	}

	authPayload := ctx.MustGet(authPayloadKey).(*token.Payload)

	arg := authPayload.UserID
	order, err := server.store.CreateOrder(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse((err)))
		return
	}

	print(len(req.OrderItems))
	for _, o := range req.OrderItems {

		println(o.ProductID)
		println(o.Quantity)
		arg := db.CreateOrderItemParams{
			OrderID:   order.OrderID,
			ProductID: o.ProductID,
			Quantity:  o.Quantity,
		}

		_, err := server.store.CreateOrderItem(ctx, arg)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse((err)))
			return
		}
	}

	ctx.JSON(http.StatusOK, order)

}
