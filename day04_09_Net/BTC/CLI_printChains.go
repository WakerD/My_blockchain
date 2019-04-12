package BTC

import (
	"fmt"
	"os"
)

func (cli *CLI) printChains(nodeID string) {
	blockchain := GetBlockchainObject(nodeID)
	if blockchain == nil {
		fmt.Println("没有区块可以打印。。")
		os.Exit(1)
	}

	defer blockchain.DB.Close()
	blockchain.PrintChains()
}
