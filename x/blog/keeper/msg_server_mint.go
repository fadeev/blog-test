package keeper

import (
	"context"

	"blog/x/blog/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	amount := sdk.NewCoins(sdk.NewCoin("foo", sdkmath.NewInt(10000)))

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, amount); err != nil {
		return nil, nil
	}

	return &types.MsgMintResponse{}, nil
}
