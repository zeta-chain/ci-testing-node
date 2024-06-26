package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeta-chain/zetacore/x/authority/types"
)

// ChainInfo queries chain info
func (k Keeper) ChainInfo(
	c context.Context,
	req *types.QueryGetChainInfoRequest,
) (*types.QueryGetChainInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	// fetch chain info
	chainInfo, found := k.GetChainInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "chain info not found")
	}

	return &types.QueryGetChainInfoResponse{ChainInfo: chainInfo}, nil
}
