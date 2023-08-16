package db

import (
	"boba/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomBobaShop(t *testing.T) BobaShop {

	name := util.RandomString(10)

	shop, err := testQueries.CreateBobaShop(context.Background(), name)

	require.NoError(t, err)

	require.NotEmpty(t, shop.ShopID)

	require.NotEmpty(t, shop)

	require.Equal(t, name, shop.Name)

	return shop
}

func TestCreateBobaShop(t *testing.T) {
	createRandomBobaShop(t)
}
