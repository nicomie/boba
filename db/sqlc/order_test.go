package db

import (
	"boba/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T) Order {
	user := createRandomAccount(t)

	arg := CreateOrderParams{
		UserID: user.ID,
		Status: util.Pending,
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, order.UserID, user.ID)
	require.Equal(t, order.Status, util.Pending)

	require.NotZero(t, order.ID)
	require.NotZero(t, order.CreatedAt)

	return order
}
func TestCreateOrder(t *testing.T) {
	createRandomOrder(t)
}
