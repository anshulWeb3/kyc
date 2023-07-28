package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"kyc/x/kyc/types"
)

func (k msgServer) CreateKyc(goCtx context.Context, msg *types.MsgCreateKyc) (*types.MsgCreateKycResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetKyc(
		ctx,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var kyc = types.Kyc{
		Creator: msg.Creator,
		Address: msg.Address,
		Value:   msg.Value,
	}

	k.SetKyc(
		ctx,
		kyc,
	)
	return &types.MsgCreateKycResponse{}, nil
}

func (k msgServer) UpdateKyc(goCtx context.Context, msg *types.MsgUpdateKyc) (*types.MsgUpdateKycResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKyc(
		ctx,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var kyc = types.Kyc{
		Creator: msg.Creator,
		Address: msg.Address,
		Value:   msg.Value,
	}

	k.SetKyc(ctx, kyc)

	return &types.MsgUpdateKycResponse{}, nil
}

func (k msgServer) DeleteKyc(goCtx context.Context, msg *types.MsgDeleteKyc) (*types.MsgDeleteKycResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKyc(
		ctx,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveKyc(
		ctx,
		msg.Address,
	)

	return &types.MsgDeleteKycResponse{}, nil
}
