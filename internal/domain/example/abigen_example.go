package example

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rom/DApp_study/internal/infrastructure/chain"
	ntf "github.com/rom/DApp_study/internal/infrastructure/contract"
)

func ExampleAbiGen(ethClientInit *chain.EthClientInit) {

	ntfAddress := "0x17067e6bf1a08E40A936dBFc6c566b9aC5DB4f60"
	contract, err := ntf.NewContracts(common.HexToAddress(ntfAddress), ethClientInit.GetClient())
	if err != nil {
		panic(err)
	}

	name, err := contract.Name(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("ntf name", name)
}
