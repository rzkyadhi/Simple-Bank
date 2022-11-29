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
