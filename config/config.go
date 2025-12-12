package config

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type (
	// Config 应用配置结构体
	Config struct {
		Sepolia Sepolia `mapstructure:"sepolia"`
	}

	// Sepolia Sepolia网络配置
	Sepolia struct {
		RPCUrl     string `mapstructure:"rpcUrl"`
		PrivateKey string `mapstructure:"privateKey"`
	}
)

// NewConfig 返回应用配置
func NewConfig() (*Config, error) {
	v := viper.New()

	// 1. 读取配置文件（可选）
	v.SetConfigName("config")
	v.AddConfigPath("./config")
	v.SetConfigType("yml")

	_ = v.ReadInConfig() // 如果不存在，不报错

	// 2. 环境变量（支持大写 + 下划线）
	v.SetEnvPrefix("") // 不设置前缀
	v.AutomaticEnv()

	_ = v.BindEnv("sepolia.rpcUrl", "SEPOLIA_RPC_URL")
	_ = v.BindEnv("sepolia.privateKey", "SEPOLIA_PRIVATE_KEY")

	// 3. 绑定命令行参数
	pflag.String("sepolia.rpcUrl", "", "Sepolia RPC URL")
	pflag.String("sepolia.privateKey", "", "Sepolia Private Key")

	pflag.Parse()
	_ = v.BindPFlags(pflag.CommandLine)

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("config unmarshal error: %w", err)
	}

	return &cfg, nil
}
