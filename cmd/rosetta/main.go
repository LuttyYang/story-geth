package main

import (
	"github.com/coinbase/rosetta-geth-sdk/configuration"
	sdkTypes "github.com/coinbase/rosetta-geth-sdk/types"
	"github.com/coinbase/rosetta-geth-sdk/utils"
	"github.com/coinbase/rosetta-sdk-go/types"
	rosettaTypes "github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rosetta/client"
	"log"
)

func main() {
	// test config
	cfg := &configuration.Configuration{
		Mode: "online",
		Network: &types.NetworkIdentifier{
			Blockchain: "Story",
			Network:    "Testnet",
		},
		GenesisBlockIdentifier: &types.BlockIdentifier{
			Hash:  params.IliadGenesisHash.Hex(),
			Index: 0,
		},
		GethURL:       "http://127.0.0.1:8546",
		RemoteGeth:    true,
		Port:          36666,
		SkipGethAdmin: false,
		GethArguments: `--config=/app/ethereum/geth.toml --gcmode=archive --graphql`,
		ChainConfig:   params.IliadChainConfig,
		RosettaCfg: configuration.RosettaConfig{
			SupportRewardTx: true,
			TraceType:       configuration.GethNativeTrace,
			Currency: &rosettaTypes.Currency{
				Symbol:   "IP",
				Decimals: 18,
			},
			TracePrefix:    "",
			FilterTokens:   false,
			TokenWhiteList: []configuration.Token{},
		},
	}

	// Create a new ethereum client by leveraging SDK functionalities
	ethereumClient, err := client.NewEthereumClient(cfg)
	if err != nil {
		log.Fatalln("cannot initialize client: %w", err)
	}

	// Bootstrap to start the Rosetta API server
	err = utils.BootStrap(cfg, sdkTypes.LoadTypes(), sdkTypes.Errors, ethereumClient)
	if err != nil {
		log.Fatalln("unable to bootstrap Rosetta server: %w", err)
	}
}
