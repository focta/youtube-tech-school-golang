package db

import (
	"context"
	"testing"
	"time"

	"github.com/focta/youtube-tech-school-golang/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {

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
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.NotZero(t, account.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	want := createRandomEntry(t, account)
	get, err := testQueries.GetEntry(context.Background(), want.ID)
	require.NoError(t, err)
	require.NotEmpty(t, get)

	require.Equal(t, want.AccountID, get.AccountID)
	require.Equal(t, want.Amount, get.Amount)
	require.Equal(t, want.AccountID, get.AccountID)
	require.WithinDuration(t, want.CreatedAt, get.CreatedAt, time.Second)
}

func TestGetListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := GetListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}
	gets, err := testQueries.GetListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, gets, 5)

	for _, get := range gets {
		require.NotEmpty(t, get)
		require.Equal(t, arg.AccountID, get.AccountID)
	}
}
