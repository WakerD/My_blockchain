package BTC

//数据库名
const DBNAME = "blockchain_%s.db"

//表名
const BLOCKTABLENAME = "blocks"

const PROTOCOL = "tcp"
const COMMANDLENGTH = 12
const NODE_VERSION = 1

//12个字节+结构体序列化的字节数组

//命令
const COMMAND_VERSION = "varsion"
const COMMAND_ADDR = "addr"
const COMMAND_BLOCK = "block"
const COMMAND_INV = "inv"
const COMMAND_GETBLOCKS = "getblocks"
const COMMAND_GETDATA = "getdata"
const COMMAND_TX = "tx"

//类型
const BLOCK_TYPE = "block"
const TX_TYPE = "tx"

//type BlockChainIterator struct {
//	//当前区块的hash
//	CurrentHash []byte
//	//数据库
//	DB *bolt.DB
//}
//
//func (bcIterator *BlockChainIterator) Next() *Block {
//	block := new(Block)
//	//1、打开数据库并读取
//	err := bcIterator.DB.View(func(tx *bolt.Tx) error {
//		//2、打开数据库
//		b := tx.Bucket([]byte(BLOCKTABLENAME))
//		if b != nil {
//			//3、根据当前hash获取数据并反序列化
//			blockBytes := b.Get(bcIterator.CurrentHash)
//			block = DeserializeBlock(blockBytes)
//			//4、更新当前的hash
//			bcIterator.CurrentHash = block.PrevBlockHash
//		}
//		return nil
//	})
//	if err != nil {
//		log.Panic(err)
//	}
//	return block
//}
