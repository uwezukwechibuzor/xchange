package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgParseSecTrade = "parse_sec_trade"

var _ sdk.Msg = &MsgParseSecTrade{}

func NewMsgParseSecTrade(creator string, tradConfId string, tradConfDtTm string, rltdOrdrId string, rltdCshMvmntId string) *MsgParseSecTrade {
	return &MsgParseSecTrade{
		Creator:        creator,
		TradConfId:     tradConfId,
		TradConfDtTm:   tradConfDtTm,
		RltdOrdrId:     rltdOrdrId,
		RltdCshMvmntId: rltdCshMvmntId,
	}
}

func (msg *MsgParseSecTrade) Route() string {
	return RouterKey
}

func (msg *MsgParseSecTrade) Type() string {
	return TypeMsgParseSecTrade
}

func (msg *MsgParseSecTrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgParseSecTrade) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgParseSecTrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
