package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d"))
	fmt.Printf("valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d"))

	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	//waitCheck := "0xe41d2489571d322189246dafa5ebde1f4699f498"
	waitCheck := "0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4"
	address := common.HexToAddress(waitCheck)
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	fmt.Println("is contract", isContract)

}
