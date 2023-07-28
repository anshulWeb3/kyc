package keeper

import (
	"kyc/x/kyc/types"
)

var _ types.QueryServer = Keeper{}
