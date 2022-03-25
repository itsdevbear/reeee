package types

import (
	bongatypes "github.com/berachain/go-bonga/x/bonga/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	uint256 "github.com/holiman/uint256"
)

type VaultI interface {
	Address() sdk.AccAddress
	Deposit(uint256.Int, sdk.AccAddress) uint256.Int
	Mint(uint256.Int, sdk.AccAddress) uint256.Int
	Withdraw(uint256.Int, sdk.AccAddress, sdk.AccAddress) uint256.Int
	Redeem(uint256.Int, sdk.AccAddress, sdk.AccAddress) uint256.Int
	ConvertToShares(uint256.Int) uint256.Int
	ConvertToAssets(uint256.Int) uint256.Int
	PreviewDeposit(uint256.Int) uint256.Int
	PreviewMint(uint256.Int) uint256.Int
	PreviewWithdraw(uint256.Int) uint256.Int
	PreviewRedeem(uint256.Int) uint256.Int
}

// ValidatorSet expected properties for the set of all validators (noalias)
type ValidatorVaultSet interface {
	IterateVaults(sdk.Context, bongatypes.ValidatorI,
		func(index int64, vault VaultI) (stop bool))

	// iterate through the consensus validator set of the last block by operator address, execute func for each validator
	IterateLastVaults(sdk.Context, bongatypes.ValidatorI,
		func(index int64, vault VaultI) (stop bool))

	Vault(sdk.Context, sdk.AccAddress) VaultI // get a particular vault by address
}
