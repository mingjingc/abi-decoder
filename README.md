 ## ABI-DECODER
 decode ethereum transaction data of contract call. Inspired by [abi-decoder](https://github.com/ConsenSys/abi-decoder)

 ## Example
 ```go
var myContractAbi = `
[
	{
		"constant": true,
		"inputs": [],
		"name": "name",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_name",
				"type": "string"
			},
			{
				"name": "_age",
				"type": "uint256"
			},
			{
				"name": "_addr",
				"type": "address"
			}
		],
		"name": "save",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "age",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "addr",
		"outputs": [
			{
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]
	`

func main()  {
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
	}
}
 ```


## Contributing
- Fork this repository
- Clone your repository
- Install dependencies
- Checkout a feature branch
- Feel free to add your features
- Make sure your features are fully tested
- Open a pull request, and enjoy (:

## LICENSE
Package abi-decoder is licensed under the [MIT](/LICENSE) License.