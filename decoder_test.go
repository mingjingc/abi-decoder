package decoder

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

/*
// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

contract HelloWorld {
    string public name;
    uint256 public age;
    address public addr;
    uint256 callTimes;

    event StudentAdded(string indexed name, uint256 indexed  age, address _addr);
    function save(string memory _name, uint256 _age, address _addr) external {
        name = _name;
        age = _age;
        addr = _addr;

        emit StudentAdded(_name, _age, _addr);
    }

    function food(bytes calldata data) external  {
        data;
        callTimes++;
    }

    function food2(bytes[][][] calldata data) external  {
        data;
        callTimes++;
    }
}
*/

var testAbi = `
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
						"indexed": false,
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

func TestABIDecoder_DecodeMethod(t *testing.T) {
	var testTransactionData = "0x19e7a9660000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000005a4728ca063b522c0b728f8000000000000000000000000000000000839c6f5a014cbfa319e8fdfa01aac186638945a80000000000000000000000000000000000000000000000000000000000000006e5b08fe6988e0000000000000000000000000000000000000000000000000000"

	abiDecoder := NewABIDecoder()
	abiDecoder.SetABI(testAbi)

	md, err := abiDecoder.DecodeMethod(testTransactionData)
	if err != nil {
		t.Error(err)
	}

	age, _ := big.NewInt(0).SetString("120000000000000000000000000000000000000", 10)
	expectMd := MethodData{
		Name: "save",
		Params: []Param{
			{
				Name:  "_name",
				Value: "小明",
				Type:  "string",
			},
			{
				Name:  "_age",
				Value: age,
				Type:  "uint256",
			},
			{
				Name:  "_addr",
				Value: common.HexToAddress("0x839C6f5a014cbfA319e8fDFA01AaC186638945A8"),
				Type:  "address",
			},
		},
	}

	assert.EqualValues(t, md.Name, expectMd.Name)
	assert.True(t, len(md.Params) == len(expectMd.Params))

	for i := 0; i < len(md.Params); i++ {
		assert.Equal(t, md.Params[i], expectMd.Params[i])
	}
}
