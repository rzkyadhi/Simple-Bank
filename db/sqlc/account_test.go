package db

import (
	"context"
	"testing"
	"time"

	"github.com/rzkyadhi/Simple-Bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accountTest := createRandomAccount(t)
	accountGet, err := testQueries.GetAccount(context.Background(), accountTest.ID)

	require.NoError(t, err)
	require.NotEmpty(t, accountGet)

	require.Equal(t, accountTest.ID, accountGet.ID)
	require.Equal(t, accountTest.Owner, accountGet.Owner)
	require.Equal(t, accountTest.Balance, accountGet.Balance)
	require.Equal(t, accountTest.Currency, accountGet.Currency)
	require.WithinDuration(t, accountTest.CreatedAt, accountGet.CreatedAt, time.Second)
}
