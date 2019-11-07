package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		networkId, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		from := ""
		msg, err := tx.AsMessage(types.NewEIP155Signer(networkId))
		if err != nil {
			from = msg.From().Hex()
		}

		fmt.Println(from, msg.From().Hex(), tx.Hash().Hex(), tx.To().Hex(), tx.Value())
	}
}
