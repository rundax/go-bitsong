package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	DefaultCodespace      = ModuleName
	ErrUnknownContract    = sdkerrors.Register(ModuleName, 1, "unknown contract")
	ErrDuplicatedContract = sdkerrors.Register(ModuleName, 4, "contractID is duplicated")
)
