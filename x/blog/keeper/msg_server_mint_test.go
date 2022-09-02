package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"blog/x/blog/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMint(t *testing.T) {
	creator := "cosmos1k86q92knalsqd2535dvk0cp5p5t2dq0tk7tv83"

	for _, tc := range []struct {
		desc    string
		request *types.MsgMint
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgMint{
				From:    creator,
				Address: creator,
				Amount:  sdk.NewCoin("foo", sdkmath.NewInt(10000))},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.Mint(ctx, tc.request)
			require.NoError(t, err)
		})
	}
}
