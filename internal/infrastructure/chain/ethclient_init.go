package chain

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClientInit 负责以太坊客户端的初始化和连接管理
type EthClientInit struct {
	client *ethclient.Client
}

// NewEthClientInit 创建一个新的以太坊客户端连接
func NewEthClientInit(rpcURL string) (*EthClientInit, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %w", err)
	}

	// 检查连接是否成功
	_, err = client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	log.Println("Successfully connected to Sepolia network")

	return &EthClientInit{
		client: client,
	}, nil
}

// GetClient 返回底层的ethclient实例
func (e *EthClientInit) GetClient() *ethclient.Client {
	return e.client
}

// Close 关闭与以太坊客户端的连接
func (e *EthClientInit) Close() {
	if e.client != nil {
		e.client.Close()
	}
}
