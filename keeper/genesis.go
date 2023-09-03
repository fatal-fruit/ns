package keeper

import (
	"context"
	nstypes "github.com/fatal-fruit/ns/types"
)

func (k *Keeper) InitGenesis(ctx context.Context, data *nstypes.GenesisState) error {
	// TODO: Implement
	return nil
}

func (k *Keeper) ExportGenesis(ctx context.Context) (*nstypes.GenesisState, error) {
	// TODO: Implement
	return &nstypes.GenesisState{}, nil
}
