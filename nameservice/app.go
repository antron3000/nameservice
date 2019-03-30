package app

import (
  "encoding/json"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/antron3000/cosmos-sdk/codec"
	"github.com/antron3000/cosmos-sdk/x/auth"
	"github.com/antron3000/cosmos-sdk/x/bank"
	"github.com/antron3000/cosmos-sdk/x/params"
	"github.com/antron3000/cosmos-sdk/x/staking"
	"github.com/antron3000/sdk-application-tutorial/x/nameservice"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	tmtypes "github.com/tendermint/tendermint/types"

  "github.com/tendermint/tendermint/libs/log"
  "github.com/cosmos/cosmos-sdk/x/auth"


)


const (
    appName = "nameservice"
)

type nameServiceApp struct {
    *bam.BaseApp
}

func NewNameServiceApp(logger log.Logger, db dbm.DB) *nameServiceApp {

    // First define the top level codec that will be shared by the different modules. Note: Codec will be explained later
    cdc := MakeCodec()

    // BaseApp handles interactions with Tendermint through the ABCI protocol
    bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

    var app = &nameServiceApp{
        BaseApp: bApp,
        cdc:     cdc,
    }

    return app
}
