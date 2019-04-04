package BTC

func (cli *CLI) createWallet() {
	wallets := NewWallets()
	wallets.CreateNewWallet()
}
