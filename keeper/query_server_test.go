package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	nstypes "github.com/fatal-fruit/ns/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryParams(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	nameRes := &nstypes.MsgBid{
		"bob.cosmos",
		f.addrs[0].String(),
		f.addrs[1].String(),
		sdk.NewCoins(sdk.NewInt64Coin("stake", 150)),
	}

	nameReq := &nstypes.QueryNameRequest{
		"bob.cosmos",
	}

	_, err := f.msgServer.Bid(f.ctx, nameRes)
	require.NoError(err)
	resp, err := f.queryServer.Name(f.ctx, nameReq)
	require.NoError(err)
	require.Equal(nameReq, resp)
}
