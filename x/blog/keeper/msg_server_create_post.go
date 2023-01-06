package keeper

import (
	"context"
	"strconv"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

	err := storedPost.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SetStoredPost(ctx, storedPost)
	postCount.Count++
	k.Keeper.SetPostCount(ctx, postCount)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.PostCreatedEventType,
			sdk.NewAttribute(types.PostCreatedCreator, msg.Creator),
			sdk.NewAttribute(types.PostCreatedPostIndex, newIndex),
			sdk.NewAttribute(types.PostCreatedTitle, msg.Title),
			sdk.NewAttribute(types.PostCreatedBody, msg.Body),
		),
	)

	return &types.MsgCreatePostResponse{
		PostIndex: newIndex,
	}, nil
}
