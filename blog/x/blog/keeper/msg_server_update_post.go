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

	post, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	creator := post.Creator
	editorAdd := msg.Creator
	if editorAdd != creator && !isEditor(post.EditorPublicKeys, editorAdd) {
		fmt.Println(types.Red+"You are not allowed to delete this post! "+types.Reset, creator)
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "You are not allowed to edit this post!")
	}
	updatedPost := types.Post{
		Creator:          creator,
		Id:               msg.Id,
		Title:            msg.Title,
		Body:             msg.Body,
		LastUpdatedAt:    ctx.BlockTime().UTC().Unix(),
		CreatedAt:        post.CreatedAt,
		EditorPublicKeys: post.EditorPublicKeys,
	}
	//todo we have a problem here, id zero don't show up in list (omit id == 0)
	k.SetPost(ctx, updatedPost)
	fmt.Println(types.BrightMagenta+"Post updated successfully"+types.Reset, msg.Creator)
	return &types.MsgUpdatePostResponse{}, nil
}

func isEditor(allowedList []string, address string) bool {
	if address == "" {
		return false
	}

	for _, str := range allowedList {
		if str == address {
			return true
		}
	}
	return false
}
