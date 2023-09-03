package keeper_test

import (
	nstypes "github.com/fatal-fruit/ns/types"
	"go.uber.org/mock/gomock"
	"testing"

	storetypes "cosmossdk.io/store/types"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/fatal-fruit/ns/keeper"
	nstestutil "github.com/fatal-fruit/ns/testutil"
)

type testFixture struct {
	ctx         sdk.Context
	k           keeper.Keeper
	msgServer   nstypes.MsgServer
	queryServer nstypes.QueryServer

	addrs []sdk.AccAddress
}

func initFixture(t *testing.T) *testFixture {
	encCfg := moduletestutil.MakeTestEncodingConfig()
	key := storetypes.NewKVStoreKey(nstypes.ModuleName)
	testCtx := testutil.DefaultContextWithDB(t, key, storetypes.NewTransientStoreKey("transient_test"))
	storeService := runtime.NewKVStoreService(key)
	addrs := simtestutil.CreateIncrementalAccounts(3)
	// gomock initializations
	ctrl := gomock.NewController(t)
	bankKeeper := nstestutil.NewMockBankKeeper(ctrl)
	defaultDenom := "stake"
	k := keeper.NewKeeper(
		encCfg.Codec,
		addresscodec.NewBech32Codec("cosmos"),
		storeService,
		addrs[0].String(),
		bankKeeper,
		defaultDenom,
	)

	err := k.InitGenesis(testCtx.Ctx, nstypes.NewGenesisState())
	if err != nil {
		panic(err)
	}

	return &testFixture{
		ctx:         testCtx.Ctx,
		k:           k,
		msgServer:   keeper.NewMsgServerImpl(k),
		queryServer: keeper.NewQueryServerImpl(k),
		addrs:       addrs,
	}
}
