package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	bongatypes "github.com/berachain/go-bonga/x/bonga/types"
)

var (
	coin100 = sdk.NewInt64Coin("steak", 100)
	coin50  = sdk.NewInt64Coin("steak", 50)
	delAddr = sdk.AccAddress("_____delegator _____")
	val1    = sdk.ValAddress("_____validator1_____")
	val2    = sdk.ValAddress("_____validator2_____")
	val3    = sdk.ValAddress("_____validator3_____")
)

func TestAuthzAuthorizations(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// verify ValidateBasic returns error for the AUTHORIZATION_TYPE_UNSPECIFIED authorization type
	delAuth, err := bongatypes.NewStakeAuthorization([]sdk.ValAddress{val1, val2}, []sdk.ValAddress{}, bongatypes.AuthorizationType_AUTHORIZATION_TYPE_UNSPECIFIED, &coin100)
	require.NoError(t, err)
	require.Error(t, delAuth.ValidateBasic())

	// verify MethodName
	delAuth, err = bongatypes.NewStakeAuthorization([]sdk.ValAddress{val1, val2}, []sdk.ValAddress{}, bongatypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE, &coin100)
	require.NoError(t, err)
	require.Equal(t, delAuth.MsgTypeURL(), sdk.MsgTypeURL(&bongatypes.MsgDelegate{}))

	// error both allow & deny list
	_, err = bongatypes.NewStakeAuthorization([]sdk.ValAddress{val1, val2}, []sdk.ValAddress{val1}, bongatypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE, &coin100)
	require.Error(t, err)

	// verify MethodName
	undelAuth, _ := bongatypes.NewStakeAuthorization([]sdk.ValAddress{val1, val2}, []sdk.ValAddress{}, bongatypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE, &coin100)
	require.Equal(t, undelAuth.MsgTypeURL(), sdk.MsgTypeURL(&bongatypes.MsgUndelegate{}))

	// verify MethodName
	beginRedelAuth, _ := bongatypes.NewStakeAuthorization([]sdk.ValAddress{val1, val2}, []sdk.ValAddress{}, bongatypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE, &coin100)
	require.Equal(t, beginRedelAuth.MsgTypeURL(), sdk.MsgTypeURL(&bongatypes.MsgBeginRedelegate{}))

	validators1_2 := []string{val1.String(), val2.String()}

	testCases := []struct {
		msg                  string
		allowed              []sdk.ValAddress
		denied               []sdk.ValAddress
		msgType              bongatypes.AuthorizationType
		limit                *sdk.Coin
		srvMsg               sdk.Msg
		expectErr            bool
		isDelete             bool
		updatedAuthorization *bongatypes.StakeAuthorization
	}{
		{
			"delegate: expect 0 remaining coins",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE,
			&coin100,
			bongatypes.NewMsgDelegate(delAddr, val1, coin100),
			false,
			true,
			nil,
		},
		{
			"delegate: verify remaining coins",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE,
			&coin100,
			bongatypes.NewMsgDelegate(delAddr, val1, coin50),
			false,
			false,
			&bongatypes.StakeAuthorization{
				Validators: &bongatypes.StakeAuthorization_AllowList{
					AllowList: &bongatypes.StakeAuthorization_Validators{Address: validators1_2},
				}, MaxTokens: &coin50, AuthorizationType: bongatypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE},
		},
		{
			"delegate: testing with invalid validator",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE,
			&coin100,
			bongatypes.NewMsgDelegate(delAddr, val3, coin100),
			true,
			false,
			nil,
		},
		{
			"delegate: testing delegate without spent limit",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE,
			nil,
			bongatypes.NewMsgDelegate(delAddr, val2, coin100),
			false,
			false,
			&bongatypes.StakeAuthorization{
				Validators: &bongatypes.StakeAuthorization_AllowList{
					AllowList: &bongatypes.StakeAuthorization_Validators{Address: validators1_2},
				}, MaxTokens: nil, AuthorizationType: bongatypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE},
		},
		{
			"delegate: fail validator denied",
			[]sdk.ValAddress{},
			[]sdk.ValAddress{val1},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE,
			nil,
			bongatypes.NewMsgDelegate(delAddr, val1, coin100),
			true,
			false,
			nil,
		},

		{
			"undelegate: expect 0 remaining coins",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE,
			&coin100,
			bongatypes.NewMsgUndelegate(delAddr, val1, coin100),
			false,
			true,
			nil,
		},
		{
			"undelegate: verify remaining coins",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE,
			&coin100,
			bongatypes.NewMsgUndelegate(delAddr, val1, coin50),
			false,
			false,
			&bongatypes.StakeAuthorization{
				Validators: &bongatypes.StakeAuthorization_AllowList{
					AllowList: &bongatypes.StakeAuthorization_Validators{Address: validators1_2},
				}, MaxTokens: &coin50, AuthorizationType: bongatypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE},
		},
		{
			"undelegate: testing with invalid validator",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE,
			&coin100,
			bongatypes.NewMsgUndelegate(delAddr, val3, coin100),
			true,
			false,
			nil,
		},
		{
			"undelegate: testing delegate without spent limit",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE,
			nil,
			bongatypes.NewMsgUndelegate(delAddr, val2, coin100),
			false,
			false,
			&bongatypes.StakeAuthorization{
				Validators: &bongatypes.StakeAuthorization_AllowList{
					AllowList: &bongatypes.StakeAuthorization_Validators{Address: validators1_2},
				}, MaxTokens: nil, AuthorizationType: bongatypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE},
		},
		{
			"undelegate: fail cannot undelegate, permission denied",
			[]sdk.ValAddress{},
			[]sdk.ValAddress{val1},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE,
			&coin100,
			bongatypes.NewMsgUndelegate(delAddr, val1, coin100),
			true,
			false,
			nil,
		},

		{
			"redelegate: expect 0 remaining coins",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE,
			&coin100,
			bongatypes.NewMsgUndelegate(delAddr, val1, coin100),
			false,
			true,
			nil,
		},
		{
			"redelegate: verify remaining coins",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE,
			&coin100,
			bongatypes.NewMsgBeginRedelegate(delAddr, val1, val1, coin50),
			false,
			false,
			&bongatypes.StakeAuthorization{
				Validators: &bongatypes.StakeAuthorization_AllowList{
					AllowList: &bongatypes.StakeAuthorization_Validators{Address: validators1_2},
				}, MaxTokens: &coin50, AuthorizationType: bongatypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE},
		},
		{
			"redelegate: testing with invalid validator",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE,
			&coin100,
			bongatypes.NewMsgBeginRedelegate(delAddr, val3, val3, coin100),
			true,
			false,
			nil,
		},
		{
			"redelegate: testing delegate without spent limit",
			[]sdk.ValAddress{val1, val2},
			[]sdk.ValAddress{},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE,
			nil,
			bongatypes.NewMsgBeginRedelegate(delAddr, val2, val2, coin100),
			false,
			false,
			&bongatypes.StakeAuthorization{
				Validators: &bongatypes.StakeAuthorization_AllowList{
					AllowList: &bongatypes.StakeAuthorization_Validators{Address: validators1_2},
				}, MaxTokens: nil, AuthorizationType: bongatypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE},
		},
		{
			"redelegate: fail cannot undelegate, permission denied",
			[]sdk.ValAddress{},
			[]sdk.ValAddress{val1},
			bongatypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE,
			&coin100,
			bongatypes.NewMsgBeginRedelegate(delAddr, val1, val1, coin100),
			true,
			false,
			nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			delAuth, err := bongatypes.NewStakeAuthorization(tc.allowed, tc.denied, tc.msgType, tc.limit)
			require.NoError(t, err)
			resp, err := delAuth.Accept(ctx, tc.srvMsg)
			require.Equal(t, tc.isDelete, resp.Delete)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tc.updatedAuthorization != nil {
					require.Equal(t, tc.updatedAuthorization.String(), resp.Updated.String())
				}
			}
		})
	}
}
