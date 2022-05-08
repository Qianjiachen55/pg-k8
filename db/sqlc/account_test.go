package db

import (
	"context"
	"github.com/Qianjiachen55/pg-k8/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount (t *testing.T){
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency:  util.RandomCurrency(),
	}
	//log.Fatal(arg)
	account,err := testQueries.CreateAccount(context.Background(),arg)

	assert.Nil(t, err,"err is not null")
	assert.NotNil(t, account,"account empty")

	assert.Equal(t, arg.Owner,account.Owner,"equal owner")
	assert.Equal(t, arg.Balance,account.Balance,"balance")
	assert.Equal(t, arg.Currency,account.Currency,"currency")

	assert.NotZero(t, account.ID,"id!")
	assert.NotZero(t, account.CreatedAt,"time!")


}