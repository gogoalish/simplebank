package db

import (
	"context"
	"testing"
	"time"

	"github.com/gogoalish/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(account Account, t *testing.T) Entry {
	arg := CreateEntryParams{account.ID, util.RandomEntry(account.Balance)}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	CreateRandomEntry(account, t)
}

func TestGetEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	entry1 := CreateRandomEntry(account, t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := CreateRandomAccount(t)
	for i := 0; i < 10; i++ {
		CreateRandomEntry(account, t)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
