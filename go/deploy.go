package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	contractAddress := "0xE590f4070EF8F16691D92D26732E9D6920FBaBaE"
	privateKey := "0x3db3135d761fa1c0df3e75fd2f172deb36e77c96b6ed07caaeb35c964356c2b3"

	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		log.Fatal(err)
	}

	// 加载合约 ABI
	abiFile, err := os.ReadFile("../contracts/MyToken.json")
	// abiFile, err := ioutil.ReadFile("../build/MyToken.abi")
	if err != nil {
		log.Fatal(err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		log.Fatal(err)
	}

	// 设置发送交易的私钥
	auth, err := bind.NewTransactor(strings.NewReader(privateKey), "")
	if err != nil {
		log.Fatal(err)
	}

	// address = common.HexToAddress(contractAddress)
	// 绑定合约
	contract := bind.NewBoundContract(common.HexToAddress(contractAddress), parsedABI, nil, client, nil)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := contract.Transact(auth, "deploy", client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contract deployed at address: %s\n", tx.To().Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

}
