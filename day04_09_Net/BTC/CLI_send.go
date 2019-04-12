package BTC

import (
	"fmt"
	"strconv"
)

func (cli *CLI) send(from, to, amount []string, nodeID string, mineNow bool) {

	blockchain := GetBlockchainObject(nodeID)

	//blockchain.MineNewBlock(from, to, amount, nodeID)
	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}

	if mineNow {
		blockchain.MineNewBlock(from, to, amount, nodeID)
		//转账成功以后，需要更新一下
		utxoSet.Update()
	} else {
		//把交易发送到矿工节点去进行验证
		fmt.Println("由矿工节点处理。。。。")
		value, _ := strconv.Atoi(amount[0])
		tx := NewSimpleTransaction(from[0], to[0], int64(value), utxoSet, []*Transaction{}, nodeID)

		sendTX(knowNodes[0], tx)
	}

	//if !dbExists() {
	//	fmt.Println("数据库不存在。。。")
	//	os.Exit(1)
	//}
	//blockchain := GetBlockchainObject()
	//
	//blockchain.MineNewBlock(from, to, amount)
	//defer blockchain.DB.Close()
	//
	//utxoSet := &UTXOSet{blockchain}
	//
	////转账成功后，需要更新
	//utxoSet.Update()
}
