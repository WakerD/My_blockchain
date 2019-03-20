package BTC

//创建一个结构体UTXO，用于表示所有未花费的
type UTXO struct {
	TxID   []byte
	Index  int
	Output *TXOutput
}
