// Package app configures and runs application.
package app

import (
	"log"

	"github.com/rom/DApp_study/config"
	"github.com/rom/DApp_study/internal/domain/example"
	"github.com/rom/DApp_study/internal/domain/service"
	"github.com/rom/DApp_study/internal/infrastructure/chain"
)

// Run 创建对象并通过构造函数运行应用
func Run() { //nolint: gocyclo,cyclop,funlen,gocritic,nolintlint
	// 加载配置
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// 连接到Sepolia网络
	ethClientInit, err := chain.NewEthClientInit(cfg.Sepolia.RPCUrl)
	if err != nil {
		log.Fatalf("Failed to connect to Sepolia network: %s", err)
	}
	defer ethClientInit.Close()

	log.Printf("Connected to Sepolia network")

	ethService := service.NewEthService(ethClientInit, cfg)

	// 执行案例
	example.Show(ethService)

}
