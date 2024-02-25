package keeper

import (
	"estore/x/estore/types"
)

var _ types.QueryServer = Keeper{}
