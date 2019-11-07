package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	host := "https://mainnet.infura.io"
	//host := "http://localhost:8545"
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal(err)
	}
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	//account := common.HexToAddress("0xB85c19ab3f062332735Cc01E80c6bfBa9F5CD31F")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance, "wei")
	fmt.Println(toEth(balance), "eth")

}

func toEth(balance *big.Int) *big.Float {
	f_balance := new(big.Float)
	f_balance.SetString(balance.String())
	eth_balance := new(big.Float).Quo(f_balance, big.NewFloat(math.Pow10(18)))
	return eth_balance
}
