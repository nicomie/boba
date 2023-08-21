package db

import (
	"boba/util"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPass, err := util.HashPin(util.RandomString(8))

	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomString(8),
		HashedPassword: hashedPass,
		Name:           util.RandomName(),
	}

	account, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.NotEmpty(t, account.HashedPassword)

	require.Equal(t, arg.Name, account.Name)
	require.Equal(t, pgtype.Int8{
		Int64: 0,
		Valid: true,
	}, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
