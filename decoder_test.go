package decoder

import "testing"
import "github.com/stretchr/testify/assert"

/*
contract HelloWorld {
    string public name;
    uint256 public age;
    address public addr;

    event StudentAdded(string indexed name, uint256 age, address _addr);
    function save(string _name, uint256 _age, address _addr) external {
        name = _name;
        age = _age;
        addr = _addr;

        emit StudentAdded(_name, _age, _addr);
    }
}
*/

var testAbi = `
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
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "name",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "age",
				"type": "uint256"
			},
			{
				"indexed": false,
				"name": "_addr",
				"type": "address"
			}
		],
		"name": "StudentAdded",
		"type": "event"
	}
]
`

var testTransactionData = "0x19e7a9660000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000005a4728ca063b522c0b728f8000000000000000000000000000000000839c6f5a014cbfa319e8fdfa01aac186638945a80000000000000000000000000000000000000000000000000000000000000006e5b08fe6988e0000000000000000000000000000000000000000000000000000"

func TestABIDecoder_DecodeMethod(t *testing.T) {
	abiDecoder := NewABIDecoder()
	abiDecoder.SetABI(testAbi)

	md, err := abiDecoder.DecodeMethod(testTransactionData)
	if err!=nil {
		t.Error(err)
	}

	expectMd := MethodData  {
		Name:"save",
		Params: []Param{
			Param{
				Name:  "_name",
				Value: "小明",
				Type:  "string",
			},
			Param{
				Name:  "_age",
				Value: "120000000000000000000000000000000000000",
				Type:  "uint256",
			},
			Param{
				Name:  "_addr",
				Value: "0x839C6f5a014cbfA319e8fDFA01AaC186638945A8",
				Type:  "address",
			},
		},
	}

	assert.EqualValues(t,md.Name,expectMd.Name)
	assert.True(t,len(md.Params) == len(expectMd.Params))

	for i:=0;i<len(md.Params);i++{
		assert.Equal(t, md.Params[i], expectMd.Params[i])
	}
}