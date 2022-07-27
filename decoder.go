package decoder

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type DecodedLog struct {
	Name    string
	Params  []Param
	Address common.Address // contract address
}

type Param struct {
	Name  string
	Value string
	Type  string
}
type MethodData struct {
	Name   string
	Params []Param
}

// ABIDecoder ethereum transaction data decoder
type ABIDecoder struct {
	myabi abi.ABI
}

func NewABIDecoder() *ABIDecoder {
	return &ABIDecoder{}
}

func (d *ABIDecoder) SetABI(contractAbi string) {
	myabi, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		log.Fatal(err)
	}
	d.myabi = myabi
}

func (d *ABIDecoder) DecodeMethod(txData string) (MethodData, error) {
	if strings.HasPrefix(txData, "0x") {
		txData = txData[2:]
	}

	decodedSig, err := hex.DecodeString(txData[:8])
	if err != nil {
		log.Fatal(err)
	}

	method, err := d.myabi.MethodById(decodedSig)
	if err != nil {
		log.Fatal(err)
	}

	decodedData, err := hex.DecodeString(txData[8:])
	if err != nil {
		log.Fatal(err)
	}

	inputs, err := method.Inputs.Unpack(decodedData)
	if err != nil {
		return MethodData{}, err
	}

	nonIndexedArgs := method.Inputs.NonIndexed()

	retData := MethodData{}
	retData.Name = method.Name
	for i, input := range inputs {
		arg := nonIndexedArgs[i]
		param := Param{
			Name:  arg.Name,
			Value: fmt.Sprintf("%v", input),
			Type:  arg.Type.String(),
		}
		retData.Params = append(retData.Params, param)
	}

	return retData, nil
}

// DecodeLogs decode contract events from transaction receipt logs
// reference: https://ethereum.stackexchange.com/questions/28637/how-to-decode-log-data-in-go
func (d *ABIDecoder) DecodeLogs(logs []*types.Log) ([]DecodedLog, error) {
	decodeLogs := make([]DecodedLog, 0, len(logs))

	for _, logItem := range logs {
		decodedLog := DecodedLog{}
		decodedLog.Address = logItem.Address

		event, err := d.myabi.EventByID(logItem.Topics[0])
		if err != nil {
			return nil, err
		}
		decodedLog.Name = event.Name
		dataList, err := d.myabi.Unpack(event.Name, logItem.Data)
		if err != nil {
			return nil, err
		}

		params := make([]Param, 0, len(event.Inputs))
		topicIndex := 1 //indexed value are put in topic
		dataIndex := 0  // no indexed value are put in data
		for _, input := range event.Inputs {
			param := Param{}

			param.Name = input.Name
			param.Type = input.Type.String()
			var value interface{}
			if input.Indexed {
				value = logItem.Topics[topicIndex]
				topicIndex++
			} else {
				value = dataList[dataIndex]
				dataIndex++
			}
			param.Value = fmt.Sprintf("%v", value)

			params = append(params, param)
		}
		decodedLog.Params = params

		decodeLogs = append(decodeLogs, decodedLog)
	}

	return decodeLogs, nil
}

func (d *ABIDecoder) ABI() abi.ABI {
	return d.myabi
}
