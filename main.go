package main

import (
	types2 "arbitrum-claim/types"
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"arbitrum-claim/distributor"
	"arbitrum-claim/proxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) < 2 {

		log.Fatal(`
			Choose mode
			live = Subscribe and waiting block target & claim
			claim = execute only claim function
			transfer = execute only transfer function. transfer all $ARB to receive address
		`)
	}

	/*
	* live = Subscribe and waiting block target
	* claim = execute only claim function
	* transfer = execute only transfer function. transfer all $ARB to receive address
	 */
	mode := os.Args[1]

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}

	config := &types2.Config{
		ContractAddressArbitrum:         os.Getenv("CONTRACT_ADDRESS_ARBITRUM"),
		ContractAddressArbitrumProxy:    os.Getenv("CONTRACT_ADDRESS_ARBITRUM_PROXY"),
		ContractAddressTokenDistributor: os.Getenv("CONTRACT_ADDRESS_TOKENDISTRIBUTOR"),
		WalletPrivateKeys:               strings.Split(strings.TrimSpace(os.Getenv("WALLET_PRIVATE_KEYS")), ","),
		EthRpcHttp:                      os.Getenv("ETH_RPC_HTTP"),
		EthRpcWss:                       os.Getenv("ETH_RPC_WSS"),
		ArbRpcHttp:                      os.Getenv("ARB_RPC_HTTP"),
		ArbRpcWss:                       os.Getenv("ARB_RPC_WSS"),
		ReceiveAddress:                  os.Getenv("RECEIVE_ADDRESS"),
		TargetBlockNo:                   os.Getenv("TARGET_BLOCK_NO"),
		Model:                           mode,
	}

	// Connect to the Ethereum node
	ethClient, err := ethclient.Dial(config.EthRpcWss)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the Arbitrum node
	arbClient, err := ethclient.Dial(config.ArbRpcHttp)
	if err != nil {
		log.Fatal(err)
	}

	// Load the smart contract
	arbitrumContractProxy, err := proxy.NewArbitrumProxy(common.HexToAddress(config.ContractAddressArbitrumProxy), arbClient)
	if err != nil {
		log.Fatal(err)
	}
	distributorContract, err := distributor.NewTokenDistributor(common.HexToAddress(config.ContractAddressTokenDistributor), arbClient)
	if err != nil {
		log.Fatal(err)
	}

	arbitrumClient := &types2.ArbitrumClient{
		EthClient:             ethClient,
		ArbClient:             arbClient,
		ArbitrumContractProxy: arbitrumContractProxy,
		DistributorContract:   distributorContract,
	}

	// 开始
	Start(arbitrumClient, config)

	//err = test.Test(arbitrumClient, config)
	//if err != nil {
	//	log.Println(err.Error())
	//}
}

func Start(arbitrumClient *types2.ArbitrumClient, config *types2.Config) {

	if config.Model == "live" {
		log.Printf("Starting Mode Live...\n")
		Subscribing(arbitrumClient, config)
	} else if config.Model == "claim" {
		log.Printf("Starting Mode Claim...\n")
		for _, privateHashKey := range config.WalletPrivateKeys {
			// go routine
			go Claim(arbitrumClient, config, privateHashKey)
		}
	} else if config.Model == "transfer" {
		log.Printf("Starting Mode Transfer...\n")
		for _, privateHashKey := range config.WalletPrivateKeys {
			// go routine
			go Transfer(arbitrumClient, config, privateHashKey, common.HexToAddress(config.ReceiveAddress))
		}
	}
}

func Subscribing(arbitrumClient *types2.ArbitrumClient, config *types2.Config) {
	headers := make(chan *types.Header)
	sub, err := arbitrumClient.EthClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	// Cast block number
	targetBlockNo, err := strconv.ParseUint(config.TargetBlockNo, 10, 64)
	if err != nil {
		panic(err)
	}

	// Subscribing Mainnet to New Blocks
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := arbitrumClient.EthClient.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			// Waiting block number is 16890400
			log.Println("New block:", block.Number().Uint64())
			if block.Number().Uint64() == targetBlockNo {
				log.Printf("Starting Claim block number: %d\n", targetBlockNo)
				// Claim
				for _, v := range config.WalletPrivateKeys {
					// go routine
					go Claim(arbitrumClient, config, v)
				}

			} else if block.Number().Uint64() == targetBlockNo+1 { // retry claim
				log.Printf("Starting Claim block number: %d\n", targetBlockNo+1)
				// Claim
				for _, v := range config.WalletPrivateKeys {
					// go routine
					go Claim(arbitrumClient, config, v)
				}
			}
		}
	}
}

func Claim(arbitrumClient *types2.ArbitrumClient, config *types2.Config, privateHashKey string) error {
	// Load private key
	privateKey, err := crypto.HexToECDSA(privateHashKey)
	if err != nil {
		return err
	}

	/*
	* Get Nonce
	 */
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("Error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	/*
	* Get gas price
	 */
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return err
	// }
	/*
	* Get gas price
	 */

	// Prepare the transactions
	chainID, _ := arbitrumClient.EthClient.ChainID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}

	txOptions := bind.TransactOpts{
		From:     fromAddress,
		Signer:   auth.Signer,
		GasLimit: 300000,
		GasPrice: big.NewInt(20 * 1e9),
	}

	// Execute a state-changing function (write) in the smart contract
	tx, err := arbitrumClient.DistributorContract.Claim(&txOptions)
	if err != nil {
		return err
	}

	log.Printf("[Claim] Wallet: %s, Transaction hash: %s\n", fromAddress, tx.Hash().Hex())

	err = Transfer(arbitrumClient, config, privateHashKey, common.HexToAddress(config.ReceiveAddress))
	if err != nil {
		return err
	}

	return nil
}

func Transfer(arbitrumClient *types2.ArbitrumClient, config *types2.Config, hexkey string, receiveAddress common.Address) error {
	// Load private key
	privateKey, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		return err
	}

	/*
	* Get Nonce
	 */
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("Error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	/*
	* Get gas price
	 */
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return err
	// }
	/*
	* Get gas price
	 */

	// Prepare the transactions
	chainID, _ := arbitrumClient.EthClient.ChainID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}

	txOptions := bind.TransactOpts{
		From:     fromAddress,
		Signer:   auth.Signer,
		GasLimit: 300000,
		GasPrice: big.NewInt(20 * 1e9),
	}

	// Check balance
	balance, err := arbitrumClient.ArbitrumContractProxy.BalanceOf(nil, fromAddress)
	if err != nil {
		return err
	}

	log.Printf("[Transfer] From: %s,  To: %s, Amount: %d\n", fromAddress, receiveAddress, balance)

	if len(balance.Bits()) == 0 {
		return errors.New("blanace is zero")
	}

	// Execute a state-changing function (write) in the smart contract
	tx, err := arbitrumClient.ArbitrumContractProxy.Transfer(&txOptions, receiveAddress, balance)
	if err != nil {
		return err
	}
	log.Printf("[Transfer] Transaction hash: %s\n", tx.Hash().Hex())
	return nil
}
