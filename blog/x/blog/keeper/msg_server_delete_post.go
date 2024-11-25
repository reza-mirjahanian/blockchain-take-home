package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"blog/x/blog/types"
)

func (k msgServer) DeletePost(goCtx context.Context, msg *types.MsgDeletePost) (*types.MsgDeletePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	post, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	creator := post.Creator
	editorAdd := msg.Creator
	if editorAdd != creator && !isEditor(post.EditorPublicKeys, editorAdd) {
		fmt.Println(types.Red+"You are not allowed to delete this post! "+types.Reset, creator)
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "You are not allowed to delete this post!")
	}

	fmt.Println(types.BrightMagenta+"Post removed successfully"+types.Reset, msg.Creator)
	k.RemovePost(ctx, msg.Id)
	return &types.MsgDeletePostResponse{}, nil
}
