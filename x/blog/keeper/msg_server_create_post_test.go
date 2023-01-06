package keeper_test

import (
	keepertest "blog/testutil/keeper"
	"blog/x/blog"
	"blog/x/blog/keeper"
	"blog/x/blog/types"
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServerCreatePost(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.BlogKeeper(t)
	blog.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreatePostSuccess(t *testing.T) {
	msgServer, _, context := setupMsgServerCreatePost(t)
	createResponse, err := msgServer.CreatePost(context, &types.MsgCreatePost{
		Title: "Test",
		Body:  "This is a test",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreatePostResponse{
		PostIndex: "0",
	}, *createResponse)
}
