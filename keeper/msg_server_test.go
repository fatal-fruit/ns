package keeper_test

import (
	nstypes "github.com/fatal-fruit/ns/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMsgReserve(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	testCases := []struct {
		name         string
		request      *nstypes.MsgBid
		expectErrMsg string
	}{
		{
			name:         "Success",
			request:      &nstypes.MsgBid{},
			expectErrMsg: "",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := f.msgServer.Bid(f.ctx, tc.request)
			if tc.expectErrMsg != "" {
				require.Error(err)
				require.ErrorContains(err, tc.expectErrMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}
