package db

import (
	"context"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, args.FromAccountID)
	require.Equal(t, transfer.ToAccountID, args.ToAccountID)
	require.Equal(t, transfer.Amount, args.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer

}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	new_transfer := createRandomTransfer(t)

	fetched_transfer, err := testQueries.GetTransfer(context.Background(), new_transfer.ID)
	require.NoError(t, err)
	require.Equal(t, new_transfer.ID, fetched_transfer.ID)
	require.Equal(t, new_transfer.FromAccountID, fetched_transfer.FromAccountID)
	require.Equal(t, new_transfer.ToAccountID, fetched_transfer.ToAccountID)
	require.Equal(t, new_transfer.Amount, fetched_transfer.Amount)
	require.WithinDuration(t, new_transfer.CreatedAt, fetched_transfer.CreatedAt, time.Second)
}
