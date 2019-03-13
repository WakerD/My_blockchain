package BTC

type BlockChain struct {
	Blocks []*Block
}

func CreateBlockChainWithGenesisBlock(data string) *BlockChain {
	genesisBlock := CreateGenesisBlock(data)
	return &BlockChain{[]*Block{genesisBlock}}
}

func (bc *BlockChain) AddBlockToBlockChain(data string, height int64, prevHash []byte) {
	newBlock := NewBlock(data, prevHash, height)
	bc.Blocks = append(bc.Blocks, newBlock)
}
