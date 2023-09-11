package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(account1, account2 Account, t *testing.T) Transfer {
	arg := CreateTransferParams{account1.ID, account2.ID, account1.Balance}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.Amount, arg.Amount)
	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)

	require.NotZero(t, transfer.CreatedAt)
	require.NotZero(t, transfer.ID)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	CreateRandomTransfer(account1, account2, t)
}

func TestGetTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	transfer1 := CreateRandomTransfer(account1, account2, t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(account1, account2, t)
		CreateRandomTransfer(account2, account1, t)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
