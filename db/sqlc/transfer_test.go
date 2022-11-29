package db

import (
	"context"
	"testing"
	"time"
	"github.com/rzkyadhi/Simple-Bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T, account1 Account, account2 Account) Transfer {
	arg := CreateTransferParams {
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	CreateRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transferTest := CreateRandomTransfer(t, account1, account2)

	transferGet, err := testQueries.GetTransfer(context.Background(), transferTest.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transferGet)

	require.Equal(t, transferTest.ID, transferGet.ID)
	require.Equal(t, transferTest.FromAccountID, transferGet.FromAccountID)
	require.Equal(t, transferTest.ToAccountID, transferGet.ToAccountID)
	require.Equal(t, transferTest.Amount, transferGet.Amount)
	require.WithinDuration(t, transferTest.CreatedAt, transferGet.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t, account1, account2)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
		require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
		require.NotEmpty(t, transfer)
	}
}
