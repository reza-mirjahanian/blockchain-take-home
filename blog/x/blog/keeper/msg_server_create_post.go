package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"blog/x/blog/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	createdTime := ctx.BlockTime().UTC().Unix()

	fmt.Println(types.BrightMagenta+"Post Creator address is: "+types.Reset, msg.Creator)
	fmt.Println(types.BrightMagenta+"Post Editor addresses are: "+types.Reset, msg.EditorPublicKeys)
	// @todo: validate editor public keys! (they should be valid addresses)
	// check uniqueness of editor public keys.
	editorPublicKeys := msg.EditorPublicKeys

	post := types.Post{
		Creator:          msg.Creator,
		Title:            msg.Title,
		Body:             msg.Body,
		CreatedAt:        createdTime,
		LastUpdatedAt:    createdTime,
		EditorPublicKeys: editorPublicKeys,
	}
	id := k.AppendPost(
		ctx,
		post,
	)
	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}
