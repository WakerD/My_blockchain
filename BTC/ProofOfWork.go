package BTC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//256位Hash里面前面至少有16个零
const TargetBit = 16

type ProofOfWork struct {
	//要验证的区块
	Block *Block

	//大整数存储目标哈希值
	Target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	//左移256-bits位
	target = target.Lsh(target, 256-TargetBit)
	return &ProofOfWork{block, target}
}

//pow挖矿算法
func (pow *ProofOfWork) Run() ([]byte, int64) {
	//1、将Block的属性拼接成字节数组
	//2、生成Hash
	//3、循环判断Hash的有效性，满足条件，跳出循环结束验证。
	nonce := 0
	hashInt := new(big.Int)
	var hash [32]byte
	for {
		//获取字节数组，为什么要通过byte类型生成hash
		dataBytes := pow.prepareData(nonce)
		//对byte数组hash签名
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%d:%x", nonce, hash)
		//将hash存储到hashInt，为什么这么做？
		hashInt.SetBytes(hash[:])
		//判断hashInt是否小于Block里的target
		//com compares x and y and returns:
		//	-1 if x<y
		//	0 if x==y
		//	1 if x>y
		if pow.Target.Cmp(hashInt) == 1 {
			break
		}
		//每次挖矿nonce值增加
		nonce++
	}
	return hash[:], int64(nonce)
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			//pow中随机变量：1、父区块的hash值，2、区块数据，3、区块当前时间戳
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.TimeStamp),
			IntToHex(int64(TargetBit)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	fmt.Println("data：", data)
	return data
}

func (pow *ProofOfWork) IsValid() bool {
	hashInt := new(big.Int)
	hashInt.SetBytes(pow.Block.Hash)
	return pow.Target.Cmp(hashInt) == 1
}
