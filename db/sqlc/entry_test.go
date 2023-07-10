package db

import (
	"context"
	"database/sql"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	new_entry := createRandomEntry(t)
	fetched_entry, err := testQueries.GetEntry(context.Background(), new_entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, fetched_entry)
	require.Equal(t, new_entry.AccountID, fetched_entry.AccountID)
	require.Equal(t, new_entry.Amount, fetched_entry.Amount)
	require.WithinDuration(t, new_entry.CreatedAt, fetched_entry.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	new_entry := createRandomEntry(t)

	update_params := UpdateEntryParams{ID: new_entry.ID, Amount: util.RandomMoney()}

	updated_entry, err := testQueries.UpdateEntry(context.Background(), update_params)

	require.NoError(t, err)
	require.NotEmpty(t, updated_entry)
	require.Equal(t, new_entry.ID, updated_entry.ID)
	require.Equal(t, new_entry.AccountID, updated_entry.AccountID)
	require.Equal(t, update_params.Amount, updated_entry.Amount)
	require.WithinDuration(t, new_entry.CreatedAt, updated_entry.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	new_entry := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), new_entry.ID)

	require.NoError(t, err)

	entry, err := testQueries.GetEntry(context.Background(), new_entry.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	entries, err := testQueries.ListEntries(context.Background(), ListEntriesParams{Offset: 5, Limit: 5})
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
