package BTC

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Height        int64
	PrevBlockHash []byte
	//Data          []byte
	Txs       []*Transaction
	TimeStamp int64
	Hash      []byte
	Nonce     int64
}

func NewBlock(txs []*Transaction, provBlockHash []byte, height int64) *Block {
	//创建区块
	block := &Block{height, provBlockHash, txs, time.Now().Unix(), nil, 0}
	//block.SetHash()
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

//func (block *Block) SetHash() {
//	heightBytes := IntToHex(block.Height)
//	timeString := strconv.FormatInt(block.TimeStamp, 2)
//	timeBytes := []byte(timeString)
//	blockBytes := bytes.Join([][]byte{
//		heightBytes,
//		block.PrevBlockHash,
//		block.Data,
//		timeBytes}, []byte{})
//
//	hash := sha256.Sum256(blockBytes)
//	block.Hash = hash[:]
//}

func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(txs, make([]byte, 32, 32), 0)
}

//将区块序列化，得到一个字节数组--区块的行为，设计为方法
func (block *Block) Serilalize() []byte {
	//1、创建一个buffer
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	var reader = bytes.NewReader(blockBytes)
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

//将Txs转为[]byte
func (block *Block) HashTransactions() []byte {
	//var txHashes [][]byte
	//var txHash [32]byte
	//for _, tx := range block.Txs {
	//	txHashes = append(txHashes, tx.TxID)
	//}
	//
	//txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	//return txHash[:]

	var txs [][]byte
	for _, tx := range block.Txs {
		txs = append(txs, tx.Serialize())
	}
	mTree := NewMerkleTree(txs)
	return mTree.RootNode.Data
}
