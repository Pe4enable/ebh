package services

import (
	"github.com/BankEx/ebh/config"
	"math/big"
)

type NodeReader struct {
	conf		*config.Config
	//httpClient	*ethclient.Client
	//wsClient	*ethclient.Client
	//rpcClient	*rpc.Client
	stopListenNode chan struct{}

}

func WeiToEth(weiValueInt *big.Int) string {

	weiValueFloat := new(big.Float).SetInt(weiValueInt)
	weiDividerFloat := big.NewFloat(1000000000000000000)

	return new(big.Float).Quo(weiValueFloat, weiDividerFloat).String()
}








