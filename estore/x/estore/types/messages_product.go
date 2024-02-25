package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProduct{}

func NewMsgCreateProduct(creator string, title string, description string, price string, quantity string) *MsgCreateProduct {
	return &MsgCreateProduct{
		Creator:     creator,
		Title:       title,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
}

func (msg *MsgCreateProduct) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProduct{}

func NewMsgUpdateProduct(creator string, id uint64, title string, description string, price string, quantity string) *MsgUpdateProduct {
	return &MsgUpdateProduct{
		Id:          id,
		Creator:     creator,
		Title:       title,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
}

func (msg *MsgUpdateProduct) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteProduct{}

func NewMsgDeleteProduct(creator string, id uint64) *MsgDeleteProduct {
	return &MsgDeleteProduct{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteProduct) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
