package types

import (
	"arbitrum-claim/distributor"
	"arbitrum-claim/proxy"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	ContractAddressArbitrum         string
	ContractAddressArbitrumProxy    string
	ContractAddressTokenDistributor string
	WalletPrivateKeys               []string
	EthRpcHttp                      string
	EthRpcWss                       string
	ArbRpcHttp                      string
	ArbRpcWss                       string
	ReceiveAddress                  string
	TargetBlockNo                   string
	Model                           string
}

type ArbitrumClient struct {
	EthClient             *ethclient.Client
	ArbClient             *ethclient.Client
	DistributorContract   *distributor.TokenDistributor
	ArbitrumContractProxy *proxy.ArbitrumProxy
}
