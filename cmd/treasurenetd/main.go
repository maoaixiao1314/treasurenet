package main

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/treasurenetprotocol/treasurenet/app"
	cmdcfg "github.com/treasurenetprotocol/treasurenet/cmd/config"

	_ "github.com/treasurenetprotocol/treasurenet/config"
)

func main() {
	setupConfig()
	cmdcfg.RegisterDenoms()

	rootCmd, _ := NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}

func setupConfig() {
	// set the address prefixes
	config := sdk.GetConfig()
	fmt.Printf("config is:%+v\n", config)
	cmdcfg.SetBech32Prefixes(config)
	cmdcfg.SetBip44CoinType(config)
	config.Seal()
}
