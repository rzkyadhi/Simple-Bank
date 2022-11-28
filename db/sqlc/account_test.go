package db

import (
	"context"
	"database/sql"
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

func TestUpdateAccount(t *testing.T) {
	accountTest := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      accountTest.ID,
		Balance: util.RandomMoney(),
	}
	accountUpdate, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accountUpdate)

	require.Equal(t, accountTest.ID, accountUpdate.ID)
	require.Equal(t, accountTest.Owner, accountUpdate.Owner)
	require.Equal(t, arg.Balance, accountUpdate.Balance)
	require.Equal(t, accountTest.Currency, accountUpdate.Currency)
	require.WithinDuration(t, accountTest.CreatedAt, accountUpdate.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	accountTest := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), accountTest.ID)
	require.NoError(t, err)

	accountDelete, err := testQueries.GetAccount(context.Background(), accountTest.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountDelete)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
