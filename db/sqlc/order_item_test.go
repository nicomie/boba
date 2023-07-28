package db

import (
	"boba/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomOrderItem(t *testing.T) OrderItem {

	order := createRandomOrder(t)

	product := createRandomProduct(t)

	qty := util.RandomInt32(1, 5)

	arg := CreateOrderItemParams{
		OrderID:   order.OrderID,
		ProductID: product.ProductID,
		Quantity:  qty,
	}

	order_item, err := testQueries.CreateOrderItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order_item)

	require.Equal(t, order_item.ProductID, product.ProductID)
	require.Equal(t, order_item.Quantity, qty)

	require.NotZero(t, order_item.OrderItemID)
	require.NotZero(t, order_item.OrderID)
	require.NotZero(t, product.ProductID)
	//require.NotZero(t, order_item.CreatedAt)

	return order_item
}
func createRandomOrderItemForOrder(t *testing.T, order Order) OrderItem {

	product := createRandomProduct(t)

	qty := util.RandomInt32(1, 5)

	arg := CreateOrderItemParams{
		OrderID:   order.OrderID,
		ProductID: product.ProductID,
		Quantity:  qty,
	}

	order_item, err := testQueries.CreateOrderItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order_item)

	require.Equal(t, order_item.ProductID, product.ProductID)
	require.Equal(t, order_item.Quantity, qty)

	require.NotZero(t, order_item.OrderItemID)
	require.NotZero(t, order_item.OrderID)
	require.NotZero(t, product.ProductID)
	//require.NotZero(t, order_item.CreatedAt)

	return order_item
}
func TestCreateOrderItem(t *testing.T) {
	createRandomOrderItem(t)
}

func TestCreateOrderItemForOrder(t *testing.T) {
	order := createRandomOrder(t)
	createRandomOrderItemForOrder(t, order)
}
