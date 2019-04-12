package BTC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"log"
)

//step1:创建一个钱包
type Wallet struct {
	//1.私钥
	PrivateKey ecdsa.PrivateKey
	//2.公钥
	PublicKey []byte
}

//产生一对秘钥
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	/*
		1.通过椭圆曲线法，随机产生私钥
		2.根据私钥生成公钥
	*/
	//椭圆加密
	curve := elliptic.P256() //椭圆加密算法，得到一个椭圆曲线值，全程：SECP256k1
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	//生成公钥
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pubKey
}

//提供一个方法用于获取钱包
func NewWallet() *Wallet {
	privateKey, publicKey := newKeyPair()
	return &Wallet{privateKey, publicKey}
}

//根据一个公钥获取对应的地址
/*
将公钥sha256 1次，再160 1次
然后version+hash
*/
func (w *Wallet) GetAddress() []byte {
	//1.先将公钥进行一次hash256，一次160，得到pubKeyHash
	PubKeyHash := PubKeyHash(w.PublicKey)
	//2.添加版本号
	versioned_payload := append([]byte{version}, PubKeyHash...)
	//3.获取校验和，将pubKeyhash,两次sha256后，取前4位
	checkSumBytes := CheckSum(versioned_payload)
	full_payload := append(versioned_payload, checkSumBytes...)
	//4.Base58
	address := Base58Encode(full_payload)
	return address
}

//一次sha256，再一次ripemd160，得到publicKeyHash
func PubKeyHash(publicKey []byte) []byte {
	//1.sha256
	hasher := sha256.New()
	hasher.Write(publicKey)
	hash := hasher.Sum(nil)

	//2.ripemd160
	ripemder := ripemd160.New()
	ripemder.Write(hash)
	pubKeyHash := ripemder.Sum(nil)

	return pubKeyHash
}

const version = byte(0x00)
const addressChecksumLen = 4

//获取验证码：将公钥哈希两次sha256，取前4位，就是校验和
func CheckSum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:addressChecksumLen]
}

//判断地址是否有效
/*
根据地址，base58解码后获取byte[],获取校验和数组
*/
func IsValidForAddress(address []byte) bool {
	full_payload := Base58Decode(address)
	fmt.Println("􏶉􏴀version_public_checksumBytes:", full_payload)
	checkSumBytes := full_payload[len(full_payload)-addressChecksumLen:]
	fmt.Println("􏶉􏴀checkSumBytes  地址后面4位？􏱨:", checkSumBytes)
	versioned_payload := full_payload[:len(full_payload)-addressChecksumLen]
	fmt.Println("􏶉􏴀version_ripemd160 地址前面4位？:", versioned_payload)
	checkBytes := CheckSum(versioned_payload)
	fmt.Println("􏶉􏴀checkBytes􏱨:", checkBytes)
	if bytes.Compare(checkSumBytes, checkBytes) == 0 {
		return true
	}
	return false
}
