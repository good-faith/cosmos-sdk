package valuerenderer_test

import (
//	"context"
//	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/valuerenderer"
	"github.com/cosmos/cosmos-sdk/baseapp"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// TODO consider to get rid of these constants, they are reused
const (
	holder     = "holder"
	multiPerm  = "multiple permissions account"
	randomPerm = "random permission"
)

/*
func initBankKeeperAndContext(t *testing.T) (keeper.BaseKeeper, types.Context) {
	app := simapp.Setup(t, false)
	c := app.BaseApp.NewContext(false, tmproto.Header{})
	maccPerms := simapp.GetMaccPerms()
	appCodec := simapp.MakeTestEncodingConfig().Codec

	maccPerms[holder] = nil
	maccPerms[authtypes.Burner] = []string{authtypes.Burner}
	maccPerms[authtypes.Minter] = []string{authtypes.Minter}
	maccPerms[multiPerm] = []string{authtypes.Burner, authtypes.Minter, authtypes.Staking}
	maccPerms[randomPerm] = []string{"random"}
	authKeeper := authkeeper.NewAccountKeeper(
		appCodec, app.GetKey(banktypes.StoreKey), app.GetSubspace(banktypes.ModuleName),
		authtypes.ProtoBaseAccount, maccPerms,
	)
	blockedAddrs := make(map[string]bool)

	keeper := keeper.NewBaseKeeper(
		appCodec, app.GetKey(banktypes.StoreKey), authKeeper,
		app.GetSubspace(banktypes.ModuleName), blockedAddrs,
	)

	return keeper,c
}
*/

// TODO add more test cases
func TestFormatCoin(t *testing.T) {
	var (
		expMetadata banktypes.Metadata
	    req  *banktypes.QueryDenomMetadataRequest
	)
	
	//bk, c := initBankKeeperAndContext(t)
	app := simapp.Setup(t, false)
	c := app.BaseApp.NewContext(false, tmproto.Header{})
	ctx := types.WrapSDKContext(c)
	queryHelper := baseapp.NewQueryServerTestHelper(c, app.InterfaceRegistry())
	banktypes.RegisterQueryServer(queryHelper, app.BankKeeper)
	// TODO consider use bankKeeper instead of queryClient
	queryClient := banktypes.NewQueryClient(queryHelper)
	
	p := message.NewPrinter(language.English)
	

	// TODO add test case to convert from mregen to uregen
	tt := []struct {
		name   string
		coin   types.Coin
		malleate func()
	}{
		{
			"convert 1000000uregen to 1regen",
			types.NewCoin("uregen", types.NewInt(int64(1000000))),
			func() {
				// TODO handle multiple metadatas 
				expMetadata = banktypes.Metadata{
					Name:        "Regen",
					Symbol:      "REGEN",
					Description: "The native staking token of the Regen network.",
					DenomUnits: []*banktypes.DenomUnit{
						{
							Denom:    "uregen",
							Exponent: 0,
							Aliases:  []string{"microregen"},
						},
						{
							Denom:    "regen",
							Exponent: 6,
							Aliases:  []string{"REGEN"},
						},
					},
					Base:    "uregen",
					Display: "regen",
				}

				app.BankKeeper.SetDenomMetaData(c, expMetadata)
				req = &banktypes.QueryDenomMetadataRequest{
					Denom: expMetadata.Base,
				}
			},

		},
		/*
		{
			"convert 1000000000uregen to 1000regen",
			types.NewCoin("uregen", types.NewInt(int64(1000000000))),
			false,
		},
		{
			"convert 23000000mregen to 23000regen",
			types.NewCoin("mregen", types.NewInt(int64(23000000))),
			false,
		},
		{
			"convert 23000000mregen to 23000000000uregen",
			types.NewCoin("mregen", types.NewInt(int64(23000000))),
			false,
		},
		{
			"convert 23000000regen to 23000000000mregen",
			types.NewCoin("regen", types.NewInt(int64(23000000))),
			false,
		},
		{
			"convert 23000regen to 23000regen",
			types.NewCoin("regen", types.NewInt(int64(23000))),
			false,
		},
		*/
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.malleate()
			queryClient.DenomsMetadata()
			res, err := queryClient.DenomMetadata(ctx, req)
			require.NoError(t, err)
			require.Equal(t, res.Metadata, expMetadata)

			dvr := valuerenderer.NewDefaultValueRenderer(expMetadata)
			formatedRes, err := dvr.Format(tc.coin)
			require.NoError(t, err)

			expAmount := p.Sprintf("%d", dvr.ComputeAmount(tc.coin))
			expDenom := expMetadata.Display
			require.Equal(t, expAmount + expDenom, formatedRes)
		})
	}
}

func TestFormatDec(t *testing.T) {
	var (
		d valuerenderer.DefaultValueRenderer
	)
	// TODO add more cases and error cases

	tt := []struct {
		name   string
		input  types.Dec
		expRes string
		expErr bool
	}{
		{
			"10 thousands decimal",
			types.NewDecFromIntWithPrec(types.NewInt(1000000), 2), // 10000.000000000000000000
			"10,000.000000000000000000",
			false,
		},
		{
			"10 mil decimal",
			types.NewInt(10000000).ToDec(),
			"10,000,000.000000000000000000",
			false,
		},

		//{"invalid string input panic", "qwerty", "", true, true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := d.Format(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expRes, res)
		})
	}
}

func TestFormatInt(t *testing.T) {
	var (
		d valuerenderer.DefaultValueRenderer
	)
	// TODO add more cases and error cases
	tt := []struct {
		name   string
		input  types.Int
		expRes string
		expErr bool
	}{
		{
			"1000000",
			types.NewInt(1000000),
			"1,000,000",
			false,
		},
		{
			"100",
			types.NewInt(100),
			"100",
			false,
		},
		{
			"23232345476756",
			types.NewInt(23232345476756),
			"23,232,345,476,756",
			false,
		},

		//{"invalid string input panic", "qwerty", "", true, true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := d.Format(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expRes, res)
		})
	}
}

// TODO add more test cases
/*
func TestParseString(t *testing.T) {
	re := regexp.MustCompile(`\d+[mu]?regen`)
	dvr := valuerenderer.NewDefaultValueRenderer()

	tt := []struct {
		str           string
		satisfyRegExp bool
		expErr        bool
	}{
		{"", false, true},
		{"10regen", true, false},
		{"1,000,000", false, false},
		{"323,000,000", false, false},
		{"1mregen", true, false},
		{"500uregen", true, false},
		{"1,500,000,000regen", true, false},
		{"394,382,328uregen", true, false},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			x, err := dvr.Parse(tc.str)
			if tc.expErr {
				require.Error(t, err)
				require.Nil(t, x)
				return
			}

			if tc.satisfyRegExp {
				require.NoError(t, err)
				coin, ok := x.(types.Coin)
				require.True(t, ok)
				require.NotNil(t, coin)
				require.True(t, re.MatchString(tc.str))
			} else {
				require.NoError(t, err)
				u, ok := x.(types.Uint)
				require.True(t, ok)
				require.NotNil(t, u)
			}
		})
	}
}
*/