package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nstypes "github.com/fatal-fruit/ns/types"
)

var _ nstypes.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) nstypes.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

func (qs queryServer) Name(goCtx context.Context, r *nstypes.QueryNameRequest) (*nstypes.QueryNameResponse, error) {
	if r == nil {
		return nil, fmt.Errorf("Empty Request")
	}
	if len(r.Name) == 0 {
		return nil, nstypes.ErrEmptyName
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	record, err := qs.k.GetNameRecord(ctx, r.Name)
	if err != nil {
		return nil, err
	}

	return &nstypes.QueryNameResponse{
		Name: &record,
	}, nil
}
