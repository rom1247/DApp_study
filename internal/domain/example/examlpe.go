package example

import (
	"fmt"

	"github.com/rom/DApp_study/internal/domain/service"
)

func Show(ethService *service.EthService) {

	// 查并打印指定区块号的区块信息 包括区块的哈希、时间戳、交易数量

	blockNum := int64(9823409)
	block, err := ethService.GetBlockByNumber(blockNum)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Block %d:\n", blockNum)
	fmt.Println("Hash:", block.Hash().Hex())
	fmt.Println("Timestamp:", block.Time())
	fmt.Println("Number of Transactions:", len(block.Transactions()))

	// 发送一个交易
	to := "0x765619a5d7551bdb3a48f68425494s7f48e40b27065ff3270a8b71b96041f51a"
	amount := int64(1000000000000000000) // 1 ETH
	tx, err := ethService.SendTransaction(to, amount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 已交易成功的交易 e5eeddf22358bb8cadbf5e5814d2ccab328c21ff6b87a20ce09c9d61020f0e57
	fmt.Println("Transaction sent:", tx.Hash().Hex())

}
