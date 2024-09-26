package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	decoder "github.com/mingjingc/abi-decoder"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var myContractAbi = `
[
			{
				"anonymous": false,
				"inputs": [
					{
						"indexed": true,
						"internalType": "string",
						"name": "name",
						"type": "string"
					},
					{
						"indexed": true,
						"internalType": "uint256",
						"name": "age",
						"type": "uint256"
					},
					{
						"indexed": false,
						"internalType": "address",
						"name": "_addr",
						"type": "address"
					}
				],
				"name": "StudentAdded",
				"type": "event"
			},
			{
				"inputs": [],
				"name": "addr",
				"outputs": [
					{
						"internalType": "address",
						"name": "",
						"type": "address"
					}
				],
				"stateMutability": "view",
				"type": "function"
			},
			{
				"inputs": [],
				"name": "age",
				"outputs": [
					{
						"internalType": "uint256",
						"name": "",
						"type": "uint256"
					}
				],
				"stateMutability": "view",
				"type": "function"
			},
			{
				"inputs": [
					{
						"internalType": "bytes",
						"name": "data",
						"type": "bytes"
					}
				],
				"name": "food",
				"outputs": [],
				"stateMutability": "nonpayable",
				"type": "function"
			},
			{
				"inputs": [
					{
						"internalType": "bytes[][][]",
						"name": "data",
						"type": "bytes[][][]"
					}
				],
				"name": "food2",
				"outputs": [],
				"stateMutability": "nonpayable",
				"type": "function"
			},
			{
				"inputs": [],
				"name": "name",
				"outputs": [
					{
						"internalType": "string",
						"name": "",
						"type": "string"
					}
				],
				"stateMutability": "view",
				"type": "function"
			},
			{
				"inputs": [
					{
						"internalType": "string",
						"name": "_name",
						"type": "string"
					},
					{
						"internalType": "uint256",
						"name": "_age",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "_addr",
						"type": "address"
					}
				],
				"name": "save",
				"outputs": [],
				"stateMutability": "nonpayable",
				"type": "function"
			}
]
`

func main() {
	txData := "0x19e7a9660000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000005a4728ca063b522c0b728f8000000000000000000000000000000000839c6f5a014cbfa319e8fdfa01aac186638945a80000000000000000000000000000000000000000000000000000000000000006e5b08fe6988e0000000000000000000000000000000000000000000000000000"

	txDataDecoder := decoder.NewABIDecoder()
	txDataDecoder.SetABI(myContractAbi)
	method, err := txDataDecoder.DecodeMethod(txData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(method.Name)
	for _, param := range method.Params {
		fmt.Println(param)
		fmt.Println(reflect.TypeOf(param.Value))
	}

	client, err := ethclient.Dial("https://rpc.ankr.com/bsc_testnet_chapel")
	if err != nil {
		log.Fatal(err)
	}

	txHash := common.HexToHash("0x66f5f9b5e0a400b9bdea8284b0013d2dbb52cfc6d76e45a46524a38cb5a7256a")
	r, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	decodedLogs, err := txDataDecoder.DecodeLogs(r.Logs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(decodedLogs)
}
