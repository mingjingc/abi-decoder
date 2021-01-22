package decoder

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Param struct {
	Name  string
	Value string
	Type  string
}
type MethodData struct {
	Name   string
	Params []*Param
}

// ethereum transaction data decoder
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

func (d *ABIDecoder) DecodeMethod(txData string) (*MethodData, error) {
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
		return nil, err
	}

	nonIndexedArgs := method.Inputs.NonIndexed()

	retData := &MethodData{}
	retData.Name = method.Name
	for i, input := range inputs {
		arg := nonIndexedArgs[i]
		param := &Param{
			Name:  arg.Name,
			Value: fmt.Sprintf("%v", input),
			Type:  arg.Type.String(),
		}
		retData.Params = append(retData.Params, param)
	}

	return retData, nil
}
