package db

import (
	"boba/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPersonnel(t *testing.T) Personnel {
	hashedPin, err := util.HashPin(util.RandomString(4))

	require.NoError(t, err)

	bobaShop := createRandomBobaShop(t)

	arg := CreatePersonnelParams{
		Name:      util.RandomString(6),
		Shop:      bobaShop.ShopID,
		HashedPin: hashedPin,
	}

	personnel, err := testQueries.CreatePersonnel(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, personnel)
	require.NotEmpty(t, personnel.BadgeNumber)
	require.NotEmpty(t, personnel.Email)

	require.Equal(t, hashedPin, personnel.HashedPin)
	require.Equal(t, bobaShop.ShopID, personnel.Shop)

	return personnel
}

func TestCreatePersonnel(t *testing.T) {
	createRandomPersonnel(t)
}
