package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount := sdk.NewCoins(msg.Amount)

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, amount); err != nil {
		return nil, nil
	}

	return &types.MsgMintResponse{}, nil
}
