package db

import (
	"context"
	"database/sql"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:     util.RandomOwner(7),
		Balance:   util.RandomMoney(),
		Currencty: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currencty, account.Currencty)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	new_account := CreateRandomAccount(t)
	fetched_account, err := testQueries.GetAcount(context.Background(), new_account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetched_account)
	require.Equal(t, new_account.Owner, fetched_account.Owner)
	require.Equal(t, new_account.Balance, fetched_account.Balance)
	require.Equal(t, new_account.Currencty, fetched_account.Currencty)
	require.WithinDuration(t, new_account.CreatedAt, fetched_account.CreatedAt, time.Second)
}
func TestUpdateAccount(t *testing.T) {
	new_account := CreateRandomAccount(t)

	update_params := UpdateAccountParams{ID: new_account.ID, Balance: util.RandomMoney()}

	updated_account, err := testQueries.UpdateAccount(context.Background(), update_params)

	require.NoError(t, err)
	require.NotEmpty(t, updated_account)
	require.Equal(t, new_account.ID, updated_account.ID)
	require.Equal(t, new_account.Owner, updated_account.Owner)
	require.Equal(t, update_params.Balance, updated_account.Balance)
	require.Equal(t, new_account.Currencty, updated_account.Currencty)
	require.WithinDuration(t, new_account.CreatedAt, updated_account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	new_account := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), new_account.ID)

	require.NoError(t, err)

	account, err := testQueries.GetAcount(context.Background(), new_account.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{Offset: 5, Limit: 5})
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
