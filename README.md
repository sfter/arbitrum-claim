# Arbitrum Claim Airdrop
### Setup

#### 1. Edit.env
    CONTRACT_ADDRESS_ARBITRUM=
    CONTRACT_ADDRESS_ARBITRUM_PROXY=
    CONTRACT_ADDRESS_TOKENDISTRIBUTOR=
    WALLET_PRIVATE_KEYS=
    ETH_RPC_HTTP=
    ETH_RPC_WSS=
    ARB_RPC_HTTP=
    ARB_RPC_WSS=
    RECEIVE_ADDRESS=
    TARGET_BLOCK_NO=

#### 2. Download dependencies
```sh
go mod tidy
go mod vendor
```

#### 3. How to run and choose mode
- live = Subscribe and waiting block target & claim
- claim = Execute only claim function
- transfer = execute only transfer function. transfer all $ARB to receive address
```sh
go run main.go live
go run main.go claim
go run main.go transfer
```
