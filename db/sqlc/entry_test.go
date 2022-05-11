package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/Qianjiachen55/pgK8/util"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	accountID := sql.NullInt64{}
	accountID.Scan(account.ID)
	arg := CreateEntryParams{
		AccountID: accountID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(),arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID,entry.AccountID)
	require.Equal(t, arg.Amount,entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t,account)
}


func TestGetEntry(t *testing.T){
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t,account)
	entry2,err := testQueries.GetEntry(context.Background(),entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	//equal
	require.Equal(t, entry1.ID,entry2.ID)
	require.Equal(t, entry1.CreatedAt,entry2.CreatedAt)
	require.Equal(t, entry1.Amount,entry2.Amount)
	require.Equal(t, entry1.AccountID,entry2.AccountID)

}

func TestQueries_ListEntries(t *testing.T) {
	account := createRandomAccount(t)

	for i:=0;i<10;i++{
		createRandomEntry(t,account)
	}
	accountID := sql.NullInt64{}
	accountID.Scan(account.ID)
	arg := ListEntriesParams{
		AccountID: accountID,
		Limit:     5,
		Offset:    5,
	}

	entries,err := testQueries.ListEntries(context.Background(),arg)
	require.NoError(t, err)
	require.Len(t, entries,5)

	for _,entry := range entries{
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID,entry.AccountID)
	}


}




