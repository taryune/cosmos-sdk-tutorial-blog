package keeper

import (
	"context"
	"strconv"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	postCount, found := k.Keeper.GetPostCount(ctx)
	if !found {
		panic("postCount is not found")
	}

	newIndex := strconv.FormatUint(postCount.Count, 10)
	storedPost := types.StoredPost{
		Index: newIndex,
		Title: msg.Title,
		Body:  msg.Body,
	}

	k.Keeper.SetStoredPost(ctx, storedPost)
	postCount.Count++
	k.Keeper.SetPostCount(ctx, postCount)

	return &types.MsgCreatePostResponse{
		PostIndex: newIndex,
	}, nil
}
