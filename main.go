package main

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	tronaddress "github.com/fbsobreira/gotron-sdk/pkg/address"
)

func encodeBalanceOfParam(address string) (string, error) {
	tronAddress, err := tronaddress.Base58ToAddress(address)
	if err != nil {
		return "", fmt.Errorf("failed to convert Base58 address to hex: %v", err)
	}

	testTronAddress := tronaddress.HexToAddress("0x41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	fmt.Println("Test Tron address:", testTronAddress.String())

	address = tronAddress.Hex()
	// address = strings.TrimPrefix(address, "0x41")
	fmt.Println("Hex address:", address)

	ethAddress := common.HexToAddress(address)

	const abiJSON = `[{"name":"balanceOf","type":"function","inputs":[{"name":"","type":"address"}]}]`

	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %v", err)
	}

	data, err := parsedABI.Pack("balanceOf", ethAddress)
	if err != nil {
		return "", fmt.Errorf("failed to encode parameter: %v", err)
	}

	return common.Bytes2Hex(data[4:]), nil
}

// https://developers.tron.network/reference/triggerconstantcontract
func main() {
	address := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	encoded, err := encodeBalanceOfParam(address)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Encoded balanceOf parameter:", encoded)	// 000000000000000000000000a614f803b6fd780986a42c78ec9c7f77e6ded13c
}
