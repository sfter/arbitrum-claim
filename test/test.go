package test

import (
	"arbitrum-claim/types"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func Test(arbitrumClient *types.ArbitrumClient, config *types.Config) error {

	// Check balance
	balance, err := arbitrumClient.ArbitrumContractProxy.BalanceOf(nil, common.HexToAddress(config.ReceiveAddress))
	if err != nil {
		println(err.Error())
		return err
	}

	log.Printf("balance: %v", balance)

	return err
}
