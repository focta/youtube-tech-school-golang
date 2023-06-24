package db

import (
	"context"
	"testing"
	"time"

	"github.com/focta/youtube-tech-school-golang/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	want := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	
	get, err := testQueries.CreateTransfer(context.Background(), want)
	require.NoError(t , err)
	require.NotEmpty(t, get)

	require.Equal(t, want.FromAccountID, get.FromAccountID)
	require.Equal(t, want.ToAccountID, get.ToAccountID)
	require.Equal(t, want.Amount, get.Amount)
	require.NotZero(t, get.ID)
	require.NotZero(t, get.CreatedAt)
}

func createRandomTransfer(t *testing.T, account1 Account, account2 Account) Transfer {

	want := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	
	get, err := testQueries.CreateTransfer(context.Background(), want)
	require.NoError(t , err)
	require.NotEmpty(t, get)

	require.Equal(t, want.FromAccountID, get.FromAccountID)
	require.Equal(t, want.ToAccountID, get.ToAccountID)
	require.Equal(t, want.Amount, get.Amount)
	require.NotZero(t, get.ID)
	require.NotZero(t, get.CreatedAt)

	return get
}

func TestGetTransfer(t *testing.T) {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	want := createRandomTransfer(t, account1, account2)

	get, err := testQueries.GetTransfer(context.Background(), want.ID)
	require.NoError(t , err)
	require.NotEmpty(t, get)

	require.Equal(t, want.FromAccountID, get.FromAccountID)
	require.Equal(t, want.ToAccountID, get.ToAccountID)
	require.Equal(t, want.Amount, get.Amount)
	require.NotZero(t, get.ID)
	require.WithinDuration(t, want.CreatedAt, get.CreatedAt, time.Second)
}

func TestGetListTransfers(t *testing.T) {
	
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2)
	}

	want := GetListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	gets, err := testQueries.GetListTransfers(context.Background(), want)
	require.NoError(t , err)
	require.Len(t, gets, 5)

	for _, get := range gets {
		require.NotEmpty(t, get)
		require.True(t, get.FromAccountID == account1.ID)
	}

}