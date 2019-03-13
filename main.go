package main

import (
	"crypto/sha256"
	"fmt"
	"golangBTC/BTC"
	"math/big"
)

func main() {
	////1、测试Block
	//block := BTC.NewBlock("I am a block", make([]byte, 32, 32), 1)
	//fmt.Printf("Height:%x\n", block.Height)
	//fmt.Printf("Data:%s\n", block.Data)
	//
	////2、测试创世区块
	//genesisBlock := BTC.CreateGenesisBlock("Genesis Block..")
	//fmt.Println(genesisBlock)
	//
	////3、测试区块链
	//genesisBlockChain := BTC.CreateBlockChainWithGenesisBlock("1")
	//fmt.Println(genesisBlockChain)
	//fmt.Println(genesisBlockChain.Blocks)
	//fmt.Printf("%s\n", genesisBlockChain.Blocks[0].Data)
	//
	////4、测试添加新区块
	//fmt.Println("creat genesis block")
	//blockChain := BTC.CreateBlockChainWithGenesisBlock("Genesis Block..")
	//blockChain.AddBlockToBlockChain("Send 100RMB To A", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//blockChain.AddBlockToBlockChain("Send 200RMB To B", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//blockChain.AddBlockToBlockChain("Send 300RMB To C", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//fmt.Println(blockChain)
	//for _, block := range blockChain.Blocks {
	//	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x/n", block.Hash)
	//	fmt.Println()
	//}

	//5、检测pow
	//a.创建一个big对象
	target := big.NewInt(256)
	fmt.Printf("0x%x\n", target)

	//b.左移256-bits位
	target = target.Lsh(target, 256-BTC.TargetBit)
	fmt.Printf("0x%x\n", target)

	s1 := "HelloWorld"
	hash := sha256.Sum256([]byte(s1))
	fmt.Printf("0x%x\n", hash)

}
