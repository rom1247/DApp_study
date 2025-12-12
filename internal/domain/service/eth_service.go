package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rom/DApp_study/config"
	"github.com/rom/DApp_study/internal/infrastructure/chain"
)

// EthService 提供以太坊相关的服务
type EthService struct {
	client *ethclient.Client

	privateKey *ecdsa.PrivateKey
}

// NewEthService 创建一个新的EthService实例
func NewEthService(ethClientInit *chain.EthClientInit, config *config.Config) *EthService {

	key := strings.TrimPrefix(config.Sepolia.PrivateKey, "0x")
	toECDSA, err := crypto.HexToECDSA(key)
	if err != nil {
		panic(err)
	}

	return &EthService{
		client:     ethClientInit.GetClient(),
		privateKey: toECDSA,
	}
}

// ctx 创建一个上下文
func (s *EthService) ctx() context.Context {

	//context.Background() 这是一个全局的context对象，没有超时，没有取消等其他设置，真实业务场景是怎么样的？
	ctx := context.Background()
	return ctx
}

// GetBlockByNumber 查询指定区块号的区块信息
func (s *EthService) GetBlockByNumber(blockNum int64) (*types.Block, error) {

	return s.client.BlockByNumber(s.ctx(), big.NewInt(blockNum))
}

// GetLatestBlock 查询最新的区块信息
func (s *EthService) GetLatestBlock() (*types.Block, error) {
	return s.client.BlockByNumber(s.ctx(), nil)
}

// SendTransaction 发送交易
func (s *EthService) SendTransaction(to string, amount int64) (*types.Transaction, error) {

	// 加载私钥
	privateKey := s.privateKey

	// 获取发送者地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取发送者账户的nonce
	nonce, err := s.client.PendingNonceAt(s.ctx(), fromAddress)
	if err != nil {
		return nil, err
	}

	// 获取预估gas价格
	gasPrice, err := s.client.SuggestGasPrice(s.ctx())

	if err != nil {
		return nil, err
	}

	toAddress := common.HexToAddress(to)

	// 创建交易
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      21000,
		To:       &toAddress,
		Value:    big.NewInt(amount),
		Data:     nil,
	})

	// 获取链ID
	chainID, err := s.client.NetworkID(s.ctx())
	if err != nil {
		return nil, err
	}

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	// 发送交易
	err = s.client.SendTransaction(s.ctx(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil

}
