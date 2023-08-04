package db

import (
	"boba/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T) Order {
	user := createRandomUser(t)

	arg := user.ID

	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, order.UserID, user.ID)
	require.Equal(t, order.Status, util.Pending)

	require.NotZero(t, order.OrderID)
	require.NotZero(t, order.CreatedAt)

	return order
}
func TestCreateOrder(t *testing.T) {
	createRandomOrder(t)
}

func TestUpdateOrderEmptyOrder(t *testing.T) {
	order := createRandomOrder(t)

	var err error
	var amount int32

	createRandomOrderItemForOrder(t, order)

	amount, err = testQueries.GetOrderItemCount(context.Background(), order.OrderID)

	require.NoError(t, err)

	err = testQueries.UpdateOrder(context.Background(), UpdateOrderParams{
		OrderID:      order.OrderID,
		CurrentItems: amount,
	})

	require.NoError(t, err)

}

func TestUpdateOrderNonEmptyOrder(t *testing.T) {
	order := createRandomOrder(t)

	OrderItem := createRandomOrderItemForOrder(t, order)

	require.NotEmpty(t, OrderItem)

	var err error
	var amount int32
	amount, err = testQueries.GetOrderItemCount(context.Background(), order.OrderID)

	require.NoError(t, err)

	err = testQueries.UpdateOrder(context.Background(), UpdateOrderParams{
		OrderID:      order.OrderID,
		CurrentItems: amount,
	})

	require.NoError(t, err)
}
