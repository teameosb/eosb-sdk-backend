package ethereum

import (
	"fmt"
	"os"
)

type EthereumEosb struct {
	*Ethereum
	*EthereumEosbProtocol
}

func NewEthereumEosb(rpcURL, hybridExAddr string) *EthereumEosb {
	if rpcURL == "" {
		rpcURL = os.Getenv("HSK_BLOCKCHAIN_RPC_URL")
	}

	if rpcURL == "" {
		panic(fmt.Errorf("NewEthereumEosb need argument rpcURL"))
	}

	return &EthereumEosb{
		NewEthereum(rpcURL, hybridExAddr),
		&EthereumEosbProtocol{},
	}
}
