package db

import (
	"context"
	"simple_bank/db/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)
func createRandomEntry(t *testing.T, id int64) Entry{
	arg:= CreateEntryParams{
		AccountID: id,
		Amount: util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, id, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	return entry
}
func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account.ID)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i:=0; i<10; i++ {
		createRandomEntry(t, account.ID)
	}
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit: 5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	for _, entry := range(entries) {
		require.NotEmpty(t, entry)
	}
	require.Len(t, entries, 5)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	arg:= CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}
	entry1 := createRandomEntry(t, arg.AccountID)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}