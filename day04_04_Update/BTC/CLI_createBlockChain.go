package BTC

func (cli *CLI) createGenesisBlockchain(address string) {
	CreateBlockChainWithGenesisBlock(address)

	bc := GetBlockchainObject()
	defer bc.DB.Close()
	if bc != nil {
		utxoSet := &UTXOSet{bc}
		utxoSet.ResetUTXOSet()
	}
}
