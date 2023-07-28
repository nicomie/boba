package db

import (
	"boba/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {

	arg := CreateProductParams{
		Name:  util.RandomName(),
		Price: int32(util.RandomMoney()),
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Price, product.Price)

	require.NotZero(t, product.ProductID)

	return product
}
func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}
