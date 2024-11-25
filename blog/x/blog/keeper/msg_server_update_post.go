package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"blog/x/blog/types"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	post := types.Post{
		Creator:       msg.Creator,
		Id:            msg.Id,
		Title:         msg.Title,
		Body:          msg.Body,
		LastUpdatedAt: ctx.BlockTime().UTC().Unix(),
		CreatedAt:     val.CreatedAt,
	}
	//todo we have a problem here, id zero don't show up in list
	k.SetPost(ctx, post)
	return &types.MsgUpdatePostResponse{}, nil
}
