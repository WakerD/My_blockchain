package BTC

func (cli *CLI) createWallet(nodeID string) {
	wallets := NewWallets(nodeID)
	wallets.CreateNewWallet(nodeID)
}
