# DApp Study

这是一个使用Go语言开发的去中心化应用(DApp)学习项目，旨在帮助开发者理解如何与以太坊区块链进行交互。

## 项目结构

```
.
├── cmd/
│   └── app/           # 应用入口点
│       └── main.go
├── config/            # 配置相关文件
│   ├── config.go
│   └── config.yml
├── internal/          # 核心业务逻辑
│   ├── app/           # 应用初始化和启动逻辑
│   │   └── app.go
│   ├── domain/        # 领域层代码
│   │   ├── example/   # 示例代码
│   │   │   └── example.go
│   │   └── service/   # 核心服务
│   │       └── eth_service.go
│   └── infrastructure/ # 基础设施层
│       └── chain/      # 区块链接入相关
│           └── ethclient_init.go
├── go.mod
└── go.sum
```

## 功能特性

- 连接以太坊Sepolia测试网络
- 查询区块信息（指定区块号或最新区块）
- 发送ETH交易

## 配置说明

项目支持多种配置方式，优先级从高到低为：命令行参数 > 环境变量 > 配置文件。

### 配置项

- `SEPOLIA_RPC_URL` 或 `sepolia.rpcUrl`: Sepolia网络的RPC端点URL
- `SEPOLIA_PRIVATE_KEY` 或 `sepolia.privateKey`: 用于交易签名的私钥

### 配置示例

#### 通过配置文件 (config/config.yml)
```yaml
sepolia:
  rpcUrl: "https://sepolia.infura.io/v3/YOUR_PROJECT_ID"
  privateKey: "YOUR_PRIVATE_KEY"
```

#### 通过环境变量
```bash
export SEPOLIA_RPC_URL="https://sepolia.infura.io/v3/YOUR_PROJECT_ID"
export SEPOLIA_PRIVATE_KEY="YOUR_PRIVATE_KEY"
```

#### 通过命令行参数
```bash
go run cmd/app/main.go --sepolia.rpcUrl="YOUR_RPC_URL" --sepolia.privateKey="YOUR_PRIVATE_KEY"
```

## 快速开始

1. 克隆项目:
   ```bash
   git clone https://github.com/rom/DApp_study.git
   cd DApp_study
   ```

2. 安装依赖:
   ```bash
   go mod tidy
   ```

3. 配置RPC URL和私钥（推荐使用环境变量）

4. 运行应用:
   ```bash
   go run cmd/app/main.go
   ```

## 核心组件说明

### EthClientInit (`internal/infrastructure/chain/ethclient_init.go`)
负责初始化和管理与以太坊节点的连接。

### EthService (`internal/domain/service/eth_service.go`)
提供核心的区块链交互功能：
- `GetBlockByNumber`: 获取指定区块号的区块信息
- `GetLatestBlock`: 获取最新区块信息
- `SendTransaction`: 发送ETH交易

### Example (`internal/domain/example/example.go`)
演示如何使用EthService进行基本的区块链操作。

## 注意事项

1. 请妥善保管您的私钥，不要将其提交到版本控制系统中
2. 在生产环境中，请使用更安全的密钥管理方案
3. 示例中的私钥仅用于测试目的