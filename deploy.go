package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// Import the generated Go bindings
)

func main() {
	// infuraURL := "https://bsc-dataseed.binance.org/"
	infuraURL := "https://data-seed-prebsc-1-s1.binance.org:8545/"
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyHex := "568f69bbf5c235d1ad74ea361ceca8f7f4bbee724a98894ab0a20f5a01d8eeae"
	// privateKeyHex = privateKeyHex[2:]
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(97))
	if err != nil {
		log.Fatalf("Failed to connect chainID:%s", big.NewInt(97))
	}

	// 1. 加载合约的ABI和BIN文件:
	abiStr, err := os.ReadFile("./build/MyToken.abi")
	if err != nil {
		log.Fatalf("Failed to read ABI: %v", err)
	}
	// 从ABI创建一个新的合约实例来与已部署的合约进行交互
	abiData, err := abi.JSON(strings.NewReader(string(abiStr)))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	binData, err := os.ReadFile("./build/MyToken.bin")
	if err != nil {
		log.Fatalf("Failed to read BIN: %v", err)
	}

	auth.GasLimit = uint64(5000000) // 5百万gas作为示例
	currentGasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	auth.GasPrice = currentGasPrice

	// 部署合约
	// address, tx, _, err := bind.DeployContract(auth, abiData, binStr, client)
	address, tx, _, err := bind.DeployContract(auth, abiData, binData, client, big.NewInt(1000000))
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	log.Println("Contract Address:", address.Hex())
	log.Println("Transaction Hash:", tx.Hash().Hex())

	// 使用已部署的合约进行代币交易：
	// tokenAddress := common.HexToAddress("YOUR_CONTRACT_ADDRESS")
	// instance, err := MyToken.NewMyToken(tokenAddress, client)
	// if err != nil {
	// 	log.Fatalf("Failed to find token contract: %v", err)
	// }

	// // 代币交易
	// tx, err := instance.Transfer(auth, common.HexToAddress("RECIPIENT_ADDRESS"), big.NewInt(AMOUNT))
	// if err != nil {
	// 	log.Fatalf("Failed to execute transfer: %v", err)
	// }

	// log.Println("Transfer TX Hash:", tx.Hash().Hex())

}
