package types

import "cosmossdk.io/collections"

const (
	ModuleName = "ns"
	StoreKey   = "ns"
)

var (
	NamesKey  = collections.NewPrefix(0)
	OwnersKey = collections.NewPrefix(1)
)
