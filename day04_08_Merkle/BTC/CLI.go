package BTC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//CLI结构体
type CLI struct {
}

func (cli *CLI) Run() {
	//判断命令行参数的长度
	isValidArgs()

	//1.创建flagset标签对象
	//addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)

	testCmd := flag.NewFlagSet("test", flag.ExitOnError)

	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	addressListsCmd := flag.NewFlagSet("addresslists", flag.ExitOnError)

	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	//2.设置标签后的参数
	//flagAddBlockData := addBlockCmd.String("data", "helloworld..", "交易数据")
	flagFromData := sendBlockCmd.String("from", "", "转账源地址")
	flagToData := sendBlockCmd.String("to", "", "转账目标地址")
	flagAmountData := sendBlockCmd.String("amount", "", "转账金额")

	flagCreateBlockChainData := createBlockChainCmd.String("address", "", "创世区块交易数据")
	flagGetBalanceData := getBalanceCmd.String("address", "", "要查询的某个账户的余额")

	//3.解析
	switch os.Args[1] {
	//case "addblock":
	//	err := addBlockCmd.Parse(os.Args[2:])
	//	if err != nil {
	//		log.Panic(err)
	//	}
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "createblockchain":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "addresslists":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "test":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		printUsage()
		os.Exit(1)
	}

	//if addBlockCmd.Parsed() {
	//	if *flagAddBlockData == "" {
	//		printUsage()
	//		os.Exit(1)
	//	}
	//	cli.addBlock(*flagAddBlockData)
	//}

	if sendBlockCmd.Parsed() {
		if *flagFromData == "" || *flagToData == "" || *flagAmountData == "" {
			printUsage()
			os.Exit(1)
		}
		fmt.Println(*flagFromData)
		fmt.Println(*flagToData)
		fmt.Println(*flagAmountData)

		from := JSONToArray(*flagFromData)
		to := JSONToArray(*flagToData)
		amount := JSONToArray(*flagAmountData)

		for i := 0; i < len(from); i++ {
			if !IsValidForAddress([]byte(from[i])) || !IsValidForAddress([]byte(to[i])) {
				fmt.Println("钱包地址无效")
				printUsage()
				os.Exit(1)
			}
		}

		cli.send(from, to, amount)
	}

	if printChainCmd.Parsed() {
		cli.printChains()
	}

	if createBlockChainCmd.Parsed() {
		//if *flagCreateBlockChainData == "" {
		//	printUsage()
		//	os.Exit(1)
		//}
		if !IsValidForAddress([]byte(*flagCreateBlockChainData)) {
			fmt.Println("钱包地址无效")
			printUsage()
			os.Exit(1)
		}
		fmt.Println("flag:" + *flagCreateBlockChainData)
		cli.createGenesisBlockchain(*flagCreateBlockChainData)
	}

	if getBalanceCmd.Parsed() {
		if *flagGetBalanceData == "" {
			fmt.Println("查询地址不能为空")
			printUsage()
			os.Exit(1)
		}
		cli.getBalance(*flagGetBalanceData)
	}

	if createWalletCmd.Parsed() {
		//创建钱包
		cli.createWallet()
	}

	//获取所有的钱包地址
	if addressListsCmd.Parsed() {
		cli.addressLists()
	}

	if testCmd.Parsed() {
		cli.TestMethod()
	}

}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreatewallet -- 创建钱包")
	fmt.Println("\taddresslists -- 输出所有钱包地址")
	fmt.Println("\tcreateblockchain -address DATA -- 创建创世区块")
	fmt.Println("\tsend -from From -to To -amount Amount - 交易数据")
	fmt.Println("\tprintchain -- 输出信息")
	fmt.Println("\tgetbalance -address DATA -- 查询账户余额")
	fmt.Println("\ttest -- 测试")
}

//func (cli *CLI) addBlock(data string) {
//	bc := GetBlockchainObject()
//	if bc == nil {
//		fmt.Println("没有创世区块，无法添加。。")
//		os.Exit(1)
//	}
//	defer bc.DB.Close()
//	bc.AddBlockToBlockChain(data)
//}

//func (cli *CLI) printChains() {
//	bc := GetBlockchainObject()
//	if bc == nil {
//		fmt.Println("没有区块可以打印。。")
//		os.Exit(1)
//	}
//	defer bc.DB.Close()
//	bc.PrintChains()
//}

//func (cli *CLI) createGenesisBlockchain(data string) {
//	CreateBlockChainWithGenesisBlock(data)
//}
