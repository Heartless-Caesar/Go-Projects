package db

import (
	util "bank/utils"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func _createAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	//Creates mock account and error if something goes wrong
	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	//Verifies if the values match
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	//Verifies not null values
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	_createAccount(t)
}

func TestGetOneAccount(t *testing.T) {
	account_1 := _createAccount(t)
	account_2, err := testQueries.GetAccount(context.Background(), account_1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account_2)

	require.Equal(t, account_1.Owner, account_2.Owner)
	require.Equal(t, account_1.Balance, account_2.Balance)
	require.Equal(t, account_1.Currency, account_2.Currency)
	require.WithinDuration(t, account_1.CreatedAt, account_2.CreatedAt, time.Second)
}
