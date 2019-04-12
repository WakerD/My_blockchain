package BTC

import "fmt"

func (cli *CLI) addressLists(nodeID string) {
	fmt.Println("打印所有的钱包地址。。")
	//获取
	Wallets := NewWallets(nodeID)
	for address, _ := range Wallets.WalletsMap {
		fmt.Println("address:", address)
	}
}
