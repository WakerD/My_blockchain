package main

import "golangBTC/day03_22_Address/BTC"

func main() {
	////1、测试Block
	//block := BTC.NewBlock("I am a block", make([]byte, 32, 32), 1)
	//fmt.Prjintf("Height:%x\n", block.Height)
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
	//fmt.Println("---------------")
	//fmt.Println(blockChain)
	////for _, block := range blockChain.Blocks {
	////	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	////	fmt.Printf("Data: %s\n", block.Data)
	////	fmt.Printf("Hash: %x/n", block.Hash)
	////	fmt.Println()
	////}
	////fmt.Println(blockChain)
	////pow := BTC.NewProofOfWork(blockChain.Blocks[0])
	//
	//fmt.Println("---------------")
	//for _, block := range blockChain.Blocks {
	//	pow := BTC.NewProofOfWork(block)
	//	fmt.Printf("PoW:%s\n", strconv.FormatBool(pow.IsValid()))
	//}

	//fmt.Printf("%v\n", pow.IsValid())

	////5、检测pow
	////a.创建一个big对象
	//target := big.NewInt(1)
	//fmt.Printf("0x%x\n", target)
	//
	////b.左移256-bits位
	//target = target.Lsh(target, 256-BTC.TargetBit)
	//fmt.Printf("0x%x\n", target)
	//
	//s1 := "HelloWorld"
	//hash := sha256.Sum256([]byte(s1))
	//fmt.Printf("0x%x\n", hash)

	////7、测试创世区块存入数据库
	//blockchain := BTC.CreateBlockChainWithGenesisBlock("Genesis Block..")
	//fmt.Println(blockchain)
	//defer blockchain.DB.Close()
	////8、测试新添加的区块
	//blockchain.AddBlockToBlockChain("Send 100RMB to AA")
	//blockchain.AddBlockToBlockChain("Send 200RMB to BB")
	//blockchain.AddBlockToBlockChain("Send 300RMB to CC")
	//fmt.Println(blockchain)
	//blockchain.PrintChains()

	//数据库声明
	//block := BTC.NewBlock("helloworld", make([]byte, 32, 32), 0)
	//
	//db, err := bolt.Open("my.db", 0600, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()

	//err = db.Update(func(tx *bolt.Tx) error {
	//	b, err := tx.CreateBucket([]byte("MyBycket"))
	//	if err != nil {
	//		return fmt.Errorf("create bucket:%s", err)
	//	}
	//
	//	if b != nil {
	//		err := b.Put([]byte("1"), []byte("send 100 BTC to AA"))
	//		if err != nil {
	//			log.Panic("数据存储失败。。。")
	//		}
	//	}
	//	return nil
	//
	//})
	//
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Println("数据存储完毕。。。")

	//err = db.View(func(tx *bolt.Tx) error {
	//	b := tx.Bucket([]byte("MyBycket"))
	//	if b != nil {
	//		data := b.Get([]byte("1"))
	//		fmt.Println(data)
	//		fmt.Printf("%s\n", data)
	//
	//		data2 := b.Get([]byte("11"))
	//		fmt.Println(data2)
	//		fmt.Printf("%s\n", data2)
	//	}
	//	return nil
	//})
	//if err != nil {
	//	log.Panic(err)
	//}

	//err = db.Update(func(tx *bolt.Tx) error {
	//	b := tx.Bucket([]byte("blocks"))
	//	if b == nil {
	//		b, err = tx.CreateBucket([]byte("blocks"))
	//		if err != nil {
	//			log.Panic("创建表失败")
	//		}
	//	}
	//	err = b.Put([]byte("1"), block.Serilalize())
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//	return nil
	//})
	//
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	////从数据库中读取该区块数据
	//err = db.View(func(tx *bolt.Tx) error {
	//	b := tx.Bucket([]byte("blocks"))
	//	if b != nil {
	//		data := b.Get([]byte("1"))
	//		//fmt.Println("--------------")
	//		//fmt.Printf("%s\n", data)
	//		block2 := BTC.DeserializeBlock(data)
	//		fmt.Println("--------------")
	//		fmt.Printf("%v\n", block2)
	//	}
	//	return nil
	//})

	//9.CLI操作
	cli := BTC.CLI{}
	cli.Run()

}
