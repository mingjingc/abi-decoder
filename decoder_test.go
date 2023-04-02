package decoder

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	if err != nil {
		t.Error(err)
	}

	_age, _ := new(big.Int).SetString("120000000000000000000000000000000000000", 10)
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
				Value: _age,
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

var SwapABI = "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"_factoryV2\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"factoryV3\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_positionManager\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_WETH9\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"WETH9\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"approveMax\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"approveMaxMinusOne\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"approveZeroThenMax\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"approveZeroThenMaxMinusOne\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"result\",\"internalType\":\"bytes\"}],\"name\":\"callPositionManager\",\"inputs\":[{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[],\"name\":\"checkOracleSlippage\",\"inputs\":[{\"type\":\"bytes[]\",\"name\":\"paths\",\"internalType\":\"bytes[]\"},{\"type\":\"uint128[]\",\"name\":\"amounts\",\"internalType\":\"uint128[]\"},{\"type\":\"uint24\",\"name\":\"maximumTickDivergence\",\"internalType\":\"uint24\"},{\"type\":\"uint32\",\"name\":\"secondsAgo\",\"internalType\":\"uint32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[],\"name\":\"checkOracleSlippage\",\"inputs\":[{\"type\":\"bytes\",\"name\":\"path\",\"internalType\":\"bytes\"},{\"type\":\"uint24\",\"name\":\"maximumTickDivergence\",\"internalType\":\"uint24\"},{\"type\":\"uint32\",\"name\":\"secondsAgo\",\"internalType\":\"uint32\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"}],\"name\":\"exactInput\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIV3SwapRouter.ExactInputParams\",\"components\":[{\"type\":\"bytes\",\"name\":\"path\",\"internalType\":\"bytes\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountOutMinimum\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"}],\"name\":\"exactInputSingle\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIV3SwapRouter.ExactInputSingleParams\",\"components\":[{\"type\":\"address\",\"name\":\"tokenIn\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"tokenOut\",\"internalType\":\"address\"},{\"type\":\"uint24\",\"name\":\"fee\",\"internalType\":\"uint24\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountOutMinimum\",\"internalType\":\"uint256\"},{\"type\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"internalType\":\"uint160\"}]}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"}],\"name\":\"exactOutput\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIV3SwapRouter.ExactOutputParams\",\"components\":[{\"type\":\"bytes\",\"name\":\"path\",\"internalType\":\"bytes\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountInMaximum\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"}],\"name\":\"exactOutputSingle\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIV3SwapRouter.ExactOutputSingleParams\",\"components\":[{\"type\":\"address\",\"name\":\"tokenIn\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"tokenOut\",\"internalType\":\"address\"},{\"type\":\"uint24\",\"name\":\"fee\",\"internalType\":\"uint24\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountInMaximum\",\"internalType\":\"uint256\"},{\"type\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"internalType\":\"uint160\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"factory\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"factoryV2\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"enumIApproveAndCall.ApprovalType\"}],\"name\":\"getApprovalType\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"result\",\"internalType\":\"bytes\"}],\"name\":\"increaseLiquidity\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIApproveAndCall.IncreaseLiquidityParams\",\"components\":[{\"type\":\"address\",\"name\":\"token0\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"token1\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount0Min\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount1Min\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"result\",\"internalType\":\"bytes\"}],\"name\":\"mint\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIApproveAndCall.MintParams\",\"components\":[{\"type\":\"address\",\"name\":\"token0\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"token1\",\"internalType\":\"address\"},{\"type\":\"uint24\",\"name\":\"fee\",\"internalType\":\"uint24\"},{\"type\":\"int24\",\"name\":\"tickLower\",\"internalType\":\"int24\"},{\"type\":\"int24\",\"name\":\"tickUpper\",\"internalType\":\"int24\"},{\"type\":\"uint256\",\"name\":\"amount0Min\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount1Min\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"}]}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes[]\",\"name\":\"\",\"internalType\":\"bytes[]\"}],\"name\":\"multicall\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"previousBlockhash\",\"internalType\":\"bytes32\"},{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes[]\",\"name\":\"\",\"internalType\":\"bytes[]\"}],\"name\":\"multicall\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes[]\",\"name\":\"results\",\"internalType\":\"bytes[]\"}],\"name\":\"multicall\",\"inputs\":[{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"positionManager\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"pull\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"refundETH\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"selfPermit\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"v\",\"internalType\":\"uint8\"},{\"type\":\"bytes32\",\"name\":\"r\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"s\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"selfPermitAllowed\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"nonce\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"expiry\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"v\",\"internalType\":\"uint8\"},{\"type\":\"bytes32\",\"name\":\"r\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"s\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"selfPermitAllowedIfNecessary\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"nonce\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"expiry\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"v\",\"internalType\":\"uint8\"},{\"type\":\"bytes32\",\"name\":\"r\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"s\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"selfPermitIfNecessary\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"v\",\"internalType\":\"uint8\"},{\"type\":\"bytes32\",\"name\":\"r\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"s\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountOutMin\",\"internalType\":\"uint256\"},{\"type\":\"address[]\",\"name\":\"path\",\"internalType\":\"address[]\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountInMax\",\"internalType\":\"uint256\"},{\"type\":\"address[]\",\"name\":\"path\",\"internalType\":\"address[]\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"sweepToken\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"sweepToken\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"sweepTokenWithFee\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"feeBips\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"feeRecipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"sweepTokenWithFee\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"feeBips\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"feeRecipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"uniswapV3SwapCallback\",\"inputs\":[{\"type\":\"int256\",\"name\":\"amount0Delta\",\"internalType\":\"int256\"},{\"type\":\"int256\",\"name\":\"amount1Delta\",\"internalType\":\"int256\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"unwrapWETH9\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"unwrapWETH9\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"unwrapWETH9WithFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"feeBips\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"feeRecipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"unwrapWETH9WithFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"feeBips\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"feeRecipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"wrapETH\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}]},{\"type\":\"receive\",\"stateMutability\":\"payable\"}]"

func loop(params []Param, decoder *ABIDecoder) {
	for _, param := range params {

		log.Println("type", param.Type)
		log.Println("name", param.Name)
		log.Println("value", param.Value)

		switch param.Type {
		case "bytes[]":

			value, ok := param.Value.([][]uint8)
			if ok {
				decodeMethod, err := decoder.DecodeMethod(hex.EncodeToString(value[0]))
				if err != nil {
					log.Println(err.Error())
					return
				}
				log.Println(decodeMethod.Name)
				loop(decodeMethod.Params, decoder)
			}

		default:
		}
	}
}

/*
type IV3SwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}
*/

// test multi call
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func TestMultiCall(t *testing.T) {
	var txDataDecoder = NewABIDecoder()

	txDataDecoder.SetABI(SwapABI)
	txData := "0x5ae401dc000000000000000000000000000000000000000000000000000000006429343500000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000e404e45aaf000000000000000000000000a1ea0b2354f5a344110af2b6ad68e75545009a03000000000000000000000000a0d71b9877f44c744546d649147e3f1e70a937600000000000000000000000000000000000000000000000000000000000000bb80000000000000000000000007ed71f614b88a736a05ef76edad1200f00f61a53000000000000000000000000000000000000000000000000002386f26fc100000000000000000000000000000000000000000000000000daf1be705f5cd33cca000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

	method, err := txDataDecoder.DecodeMethod(txData)
	log.Println(method.Name)

	if err != nil {
		log.Fatal(err)
	}

	loop(method.Params, txDataDecoder)
}
