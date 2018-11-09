package services

import (
	"context"
	"github.com/BankEx/ebh/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"github.com/BankEx/ebh/models"
	"time"
	"log"
	"strconv"
	"fmt"
)

type NodeReader struct {
	conf		*config.Config
	httpClient	*ethclient.Client
	wsClient	*ethclient.Client
	rpcClient	*rpc.Client
	stopListenNode chan struct{}

}

func WeiToEth(weiValueInt *big.Int) string {

	weiValueFloat := new(big.Float).SetInt(weiValueInt)
	weiDividerFloat := big.NewFloat(1000000000000000000)

	return new(big.Float).Quo(weiValueFloat, weiDividerFloat).String()
}


func GetFromAddress(tx *types.Transaction) (string, error) {
	sgn := types.NewEIP155Signer(tx.ChainId())
	fromAddrObj, err := sgn.Sender(tx)
	if err == nil {
		return fromAddrObj.String(), nil
	}
	return "", err
}


func NewNodeReader(conf *config.Config) (service *NodeReader, err error) {
	service = new(NodeReader)
	service.conf = conf

	service.httpClient, err = ethclient.Dial("http://51.144.244.221:80")
	//service.conf.NodeConfig.Host)
	if err != nil {
		return nil, err
	}

	service.wsClient, err = ethclient.Dial("ws://51.144.244.221:60046")
	if err != nil {
		return nil, err
	}

	service.rpcClient, err = rpc.Dial("http://51.144.244.221:80")
	//http://sammy:jEKuTJHkTOOWkopnXqC1d1@51.144.244.221:80
	if err != nil {
		return nil, err
	}

	service.SubscribeNewBlock(6)

	return service, nil
}

//func (s *NodeReader) ConnectToDB() (err error) {
//	s.dbConn, err = repositories.NewDBConn(s.conf.DBConfig.URL, s.conf.DBConfig.DBName)
//	return
//}
//}

func (s *NodeReader) SubscribeNewBlock(numOfConfirmations int64) {

	heads := make(chan *types.Header, 1024)
	_, err := s.wsClient.SubscribeNewHead(context.Background(), heads)

	if err == nil {

		go func() {
			for {
				select {
				case head := <-heads:
					lastBlockNum := head.Number.Int64()
					log.Printf("New block is appear: %d", lastBlockNum)

					pastBlock := lastBlockNum - numOfConfirmations
					log.Printf("Block with %d confirmations is %d", numOfConfirmations, pastBlock)
				}
			}
		}()
	} else {
		log.Fatal("RPC WS connection is lost")
	}

}


func (s *NodeReader) GetBlockByHash(hash string) (*types.Block, error) {

	h := common.HexToHash(hash)

	result, err := s.httpClient.BlockByHash(context.Background(), h)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *NodeReader) GetBlockByNumber(number int64) ([]models.TxnHeader, error) {

	result := make([]models.TxnHeader, 0)

	blockNumber := big.NewInt(number)

	block, err := s.httpClient.BlockByNumber(context.Background(), blockNumber)
	timestamp := block.Time().Int64()
	datetime := time.Unix(timestamp, 0)

	if err != nil {
		return result, err
	}

	transactions := block.Transactions()

	for _, tx := range transactions{

		ethValue := WeiToEth(tx.Value())

		gasUsed := big.NewInt(int64(tx.Gas()))
		weiFee := new(big.Int).Mul(tx.GasPrice(), gasUsed)
		ethFee := WeiToEth(weiFee)

		fromAddr, err := GetFromAddress(tx)

		// TODO: TX status?
		// To get TX status you need request TX Receipt (s.client.TransactionReceipt()) for EACH transaction
		// Each TX Receipt is a separated request

		if err == nil {
			txHeader := models.TxnHeader{
				Hash: tx.Hash().String(),
				Block: number,
				From: fromAddr,
				To: tx.To().String(),
				Datetime: datetime,
				Value: ethValue,
				Fee: ethFee,
				Status: 1,
			}

			result = append(result, txHeader)
		} else {
			log.Printf("Error during getting from address: %s.", err)
		}
	}
	return result, nil
}

func (s *NodeReader) GetHeaderByNumber(number int64) (*types.Header, error) {

	blockNumber := big.NewInt(number)

	result, err := s.httpClient.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func (s *NodeReader) SendTransaction(from string, to string, value *big.Int, gasLimit int64, gasPrice int64, password string) (string, error) {

	type TxParams struct {
		From		string		`json:"from" bson:"from"`
		To			string		`json:"to" bson:"to"`
		Gas			string		`json:"gas" bson:"gas"`
		GasPrice	string		`json:"gasPrice" bson:"gasPrice"`
		Value		string		`json:"value" bson:"value"`
	}

	params := TxParams{
		From: from,
		To: to,
		Gas: "0x" + strconv.FormatInt(gasLimit, 16),
		GasPrice: "0x" + strconv.FormatInt(gasPrice, 16),
		Value: "0x" + fmt.Sprintf("%x", value),
	}

	var hash string

	err := s.rpcClient.Call(&hash, "personal_sendTransaction", params, password)
	if err != nil{
		return "0x", err
	}

	return hash, nil

}


func (s *NodeReader) ImportWallet(secret string, password string) (string, error) {

	var address string

	err := s.rpcClient.Call(&address, "parity_newAccountFromWallet", secret, password)
	if err != nil{
		return "0x", err
	}

	return address, nil

}


func (s *NodeReader) CreateHotWallet(passwd string) (*interface{}, error) {

	var address string

	// create new account and get address
	err := s.rpcClient.Call(&address, "personal_newAccount", passwd)
	if err == nil {

		// export secret params of new account (password is required)
		exportResult := new(interface{})
		err := s.rpcClient.Call(exportResult, "parity_exportAccount", address, passwd)
		if err == nil {

			// remove account ASAP after creation
			killResult := new(interface{})
			err := s.rpcClient.Call(&killResult, "parity_killAccount", address, passwd)
			if err == nil {
				return exportResult, err
			} else {
				log.Printf("Error during exporting wallet: %s.", err)
				return killResult, err
			}

		} else {
			log.Printf("Error during exporting wallet: %s.", err)
			return exportResult, err
		}

	} else {
		log.Printf("Error during creating new wallet: %s.", err)
		return new(interface{}), err
	}
}
