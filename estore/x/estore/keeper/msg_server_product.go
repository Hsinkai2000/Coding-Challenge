package keeper

import (
	"context"
	"fmt"

	"estore/x/estore/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateProduct(goCtx context.Context, msg *types.MsgCreateProduct) (*types.MsgCreateProductResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var product = types.Product{
		Creator:     msg.Creator,
		Title:       msg.Title,
		Description: msg.Description,
		Price:       msg.Price,
		Quantity:    msg.Quantity,
	}

	id := k.AppendProduct(
		ctx,
		product,
	)

	return &types.MsgCreateProductResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateProduct(goCtx context.Context, msg *types.MsgUpdateProduct) (*types.MsgUpdateProductResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var product = types.Product{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Title:       msg.Title,
		Description: msg.Description,
		Price:       msg.Price,
		Quantity:    msg.Quantity,
	}

	// Checks that the element exists
	val, found := k.GetProduct(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetProduct(ctx, product)

	return &types.MsgUpdateProductResponse{}, nil
}

func (k msgServer) DeleteProduct(goCtx context.Context, msg *types.MsgDeleteProduct) (*types.MsgDeleteProductResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetProduct(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveProduct(ctx, msg.Id)

	return &types.MsgDeleteProductResponse{}, nil
}
