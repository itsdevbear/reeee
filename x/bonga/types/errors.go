package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/staking module sentinel errors
//
// TODO: Many of these errors are redundant. They should be removed and replaced
// by sdkerrors.ErrInvalidRequest.
//
// REF: https://github.com/cosmos/cosmos-sdk/issues/5450
var (
	ErrEmptyValidatorAddr              = sdkerrors.Register(ModuleName, 1002, "empty validator address")
	ErrNoValidatorFound                = sdkerrors.Register(ModuleName, 1003, "validator does not exist")
	ErrValidatorOwnerExists            = sdkerrors.Register(ModuleName, 1004, "validator already exist for this operator address; must use new validator operator address")
	ErrValidatorPubKeyExists           = sdkerrors.Register(ModuleName, 1005, "validator already exist for this pubkey; must use new validator pubkey")
	ErrValidatorPubKeyTypeNotSupported = sdkerrors.Register(ModuleName, 1006, "validator pubkey type is not supported")
	ErrValidatorJailed                 = sdkerrors.Register(ModuleName, 1007, "validator for this address is currently jailed")
	ErrBadRemoveValidator              = sdkerrors.Register(ModuleName, 1008, "failed to remove validator")
	ErrCommissionNegative              = sdkerrors.Register(ModuleName, 1009, "commission must be positive")
	ErrCommissionHuge                  = sdkerrors.Register(ModuleName, 10010, "commission cannot be more than 100%")
	ErrCommissionGTMaxRate             = sdkerrors.Register(ModuleName, 10011, "commission cannot be more than the max rate")
	ErrCommissionUpdateTime            = sdkerrors.Register(ModuleName, 10012, "commission cannot be changed more than once in 24h")
	ErrCommissionChangeRateNegative    = sdkerrors.Register(ModuleName, 10013, "commission change rate must be positive")
	ErrCommissionChangeRateGTMaxRate   = sdkerrors.Register(ModuleName, 10014, "commission change rate cannot be more than the max rate")
	ErrCommissionGTMaxChangeRate       = sdkerrors.Register(ModuleName, 10015, "commission cannot be changed more than max change rate")
	ErrSelfDelegationBelowMinimum      = sdkerrors.Register(ModuleName, 10016, "validator's self delegation must be greater than their minimum self delegation")
	ErrMinSelfDelegationDecreased      = sdkerrors.Register(ModuleName, 10017, "minimum self delegation cannot be decrease")
	ErrEmptyDelegatorAddr              = sdkerrors.Register(ModuleName, 10018, "empty delegator address")
	ErrNoDelegation                    = sdkerrors.Register(ModuleName, 10019, "no delegation for (address, validator) tuple")
	ErrBadDelegatorAddr                = sdkerrors.Register(ModuleName, 10020, "delegator does not exist with address")
	ErrNoDelegatorForAddress           = sdkerrors.Register(ModuleName, 10021, "delegator does not contain delegation")
	ErrInsufficientShares              = sdkerrors.Register(ModuleName, 10022, "insufficient delegation shares")
	ErrDelegationValidatorEmpty        = sdkerrors.Register(ModuleName, 10023, "cannot delegate to an empty validator")
	ErrNotEnoughDelegationShares       = sdkerrors.Register(ModuleName, 10024, "not enough delegation shares")
	ErrNotMature                       = sdkerrors.Register(ModuleName, 10025, "entry not mature")
	ErrNoUnbondingDelegation           = sdkerrors.Register(ModuleName, 10026, "no unbonding delegation found")
	ErrMaxUnbondingDelegationEntries   = sdkerrors.Register(ModuleName, 10027, "too many unbonding delegation entries for (delegator, validator) tuple")
	ErrNoRedelegation                  = sdkerrors.Register(ModuleName, 10028, "no redelegation found")
	ErrSelfRedelegation                = sdkerrors.Register(ModuleName, 10029, "cannot redelegate to the same validator")
	ErrTinyRedelegationAmount          = sdkerrors.Register(ModuleName, 10030, "too few tokens to redelegate (truncates to zero tokens)")
	ErrBadRedelegationDst              = sdkerrors.Register(ModuleName, 10031, "redelegation destination validator not found")
	ErrTransitiveRedelegation          = sdkerrors.Register(ModuleName, 10032, "redelegation to this validator already in progress; first redelegation to this validator must complete before next redelegation")
	ErrMaxRedelegationEntries          = sdkerrors.Register(ModuleName, 10033, "too many redelegation entries for (delegator, src-validator, dst-validator) tuple")
	ErrDelegatorShareExRateInvalid     = sdkerrors.Register(ModuleName, 10034, "cannot delegate to validators with invalid (zero) ex-rate")
	ErrBothShareMsgsGiven              = sdkerrors.Register(ModuleName, 10035, "both shares amount and shares percent provided")
	ErrNeitherShareMsgsGiven           = sdkerrors.Register(ModuleName, 10036, "neither shares amount nor shares percent provided")
	ErrInvalidHistoricalInfo           = sdkerrors.Register(ModuleName, 10037, "invalid historical info")
	ErrNoHistoricalInfo                = sdkerrors.Register(ModuleName, 10038, "no historical info found")
	ErrEmptyValidatorPubKey            = sdkerrors.Register(ModuleName, 10039, "empty validator public key")
)
