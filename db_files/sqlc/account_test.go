package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountCreation(t *testing.T){
	arg := CreateAccountParams{
		Owner: "Cesa",
		Balance: 5000,
		Currency: "R$",
	}	
	
	//Creates mock account and error if something goes wrong
	account, err := testQueries.CreateAccount(context.Background(),arg) 

	require.NoError(t,err)
	require.NotEmpty(t,account)	
	
	//Verifies if the values match
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)	

	//Verifies not null values
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
