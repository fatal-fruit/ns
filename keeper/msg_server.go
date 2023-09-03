package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	nstypes "github.com/fatal-fruit/ns/types"
)

type msgServer struct {
	k Keeper
}

var _ nstypes.MsgServer = msgServer{}

func NewMsgServerImpl(keeper Keeper) nstypes.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) Reserve(goCtx context.Context, msg *nstypes.MsgReserve) (*nstypes.MsgReserveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := ms.k.checkAvailability(ctx, msg.Name); err != nil {
		return nil, err
	}

	nameReservation, err := ms.k.validateAndFormat(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err = ms.k.reserveName(ctx, nameReservation); err != nil {
		return nil, err
	}

	return &nstypes.MsgReserveResponse{}, nil
}
