package BTC

type Inv struct {
	AddrFrom string   //自己的地址
	Type     string   //类型block tx
	Items    [][]byte //hash二维数组
}
