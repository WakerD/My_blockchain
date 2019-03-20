package BTC

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

//func (bc *BlockChain) Iterator() *BlockChainIterator {
//	return &BlockChainIterator{bc.Tip, bc.DB}
//}

func (bcIterator *BlockChainIterator) Next() *Block {
	block := new(Block)
	//1、打开数据库并读取
	err := bcIterator.DB.View(func(tx *bolt.Tx) error {
		//2、打开数据库
		b := tx.Bucket([]byte(BLOCKTABLENAME))
		if b != nil {
			//3、根据当前hash获取数据并反序列化
			blockBytes := b.Get(bcIterator.CurrentHash)
			block = DeserializeBlock(blockBytes)
			//4、更新当前的hash
			bcIterator.CurrentHash = block.PrevBlockHash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}

//func (bc *BlockChain) PrintChains() {
//	//1.获取迭代器对象
//	bcIterator := bc.Iterator()
//
//	var count = 0
//	//2.循环迭代
//	for {
//		block := bcIterator.Next()
//		count++
//		fmt.Printf("第%d个区块的信息：\n", count)
//		//获取当前hash对应的数据，并进行反序列化
//		fmt.Printf("\t高度：%d\n", block.Height)
//		fmt.Printf("\t上一个区块的hash：%x\n", block.PrevBlockHash)
//		fmt.Printf("\t当前的hash：%x\n", block.Hash)
//		fmt.Printf("\t数据：%s\n", block.Data)
//		fmt.Printf("\t时间：%s\n", time.Unix(block.TimeStamp, 0).Format("2006-01-02 15:04:05"))
//		fmt.Printf("\t次数：%d\n", block.Nonce)
//
//		//3.直到父hash值为0
//		hashInt := new(big.Int)
//		hashInt.SetBytes(block.PrevBlockHash)
//		if big.NewInt(0).Cmp(hashInt) == 0 {
//			break
//		}
//	}
//}
