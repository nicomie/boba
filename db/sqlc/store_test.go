package db

import (
	"context"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	order := createRandomOrder(t)

	n := 5
	for i := 0; i < n; i++ {
		res, err := store.TransferTx(context.Background(), TransferTxParams{
			FromAccountID: acc1.ID,
			ToAccountID:   acc2.ID,
			OrderID:       1,
		})
	}

}
