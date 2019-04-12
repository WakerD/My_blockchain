package BTC

//存储节点全局变量

//localhost:3000 主节点的地址
var knowNodes = []string{"localhost:3000"}

//全局变量，节点地址
var nodeAddress string

//存储hash值
var transactionArray [][]byte
var minerAddress string
var memoryTxPool = make(map[string]*Transaction)
