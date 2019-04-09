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
	Data          []byte
	TimeStamp     int64
	Hash          []byte
	Nonce         int64
}

func NewBlock(data string, provBlockHash []byte, height int64) *Block {
	block := &Block{height, provBlockHash, []byte(data), time.Now().Unix(), nil, 0}
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

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, make([]byte, 32, 32), 0)
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
