package db

import (
	"context"
	"testing"
	"time"

	"github.com/rzkyadhi/Simple-Bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T, account Account) Entry {
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
	account := createRandomAccount(t)
	CreateRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entryTest := CreateRandomEntry(t, account)
	entryGet, err := testQueries.GetEntry(context.Background(), entryTest.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entryGet)

	require.Equal(t, entryTest.ID, entryGet.ID)
	require.Equal(t, entryTest.AccountID, entryGet.AccountID)
	require.Equal(t, entryTest.Amount, entryGet.Amount)
	require.WithinDuration(t, entryTest.CreatedAt, entryGet.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	account := createRandomAccount(t)
	entryTest := CreateRandomEntry(t, account)
	arg := UpdateEntryParams{
		ID: entryTest.ID,
		Amount: util.RandomMoney(),
	}
	entryUpdate, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entryUpdate)

	require.Equal(t, arg.ID, entryUpdate.ID)
	require.Equal(t, arg.Amount, entryUpdate.Amount)
}

func TestListEntries(t *testing.T) {
	accountTest := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t, accountTest)
	}
	arg := ListEntriesParams{
		AccountID: accountTest.ID,
		Limit: 5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.Equal(t, arg.AccountID, entry.AccountID)
		require.NotEmpty(t, entry)
	}

}
