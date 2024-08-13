package main

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	tronaddress "github.com/fbsobreira/gotron-sdk/pkg/address"
)

func encodeBalanceOfParam(address string) (string, error) {
	// Convert Base58 address to hex
	tronAddress, err := tronaddress.Base58ToAddress(address)
	if err != nil {
		return "", fmt.Errorf("failed to convert Base58 address to hex: %v", err)
	}

	// Remove the "41" prefix if present
	address = tronAddress.Hex()
	address = strings.TrimPrefix(address, "0x41")

	// Convert the address to the appropriate type
	ethAddress := common.HexToAddress(address)

	// Create a simple ABI with just the balanceOf function
	const abiJSON = `[{"name":"balanceOf","type":"function","inputs":[{"name":"","type":"address"}]}]`

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %v", err)
	}

	// Encode the parameter
	data, err := parsedABI.Pack("balanceOf", ethAddress)
	if err != nil {
		return "", fmt.Errorf("failed to encode parameter: %v", err)
	}

	// Convert to hex string and remove the function selector (first 4 bytes)
	return common.Bytes2Hex(data[4:]), nil
}

func main() {
	address := "TZ4UXDV5ZhNW7fb2AMSbgfAEZ7hWsnYS2g"
	encoded, err := encodeBalanceOfParam(address)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Encoded balanceOf parameter:", encoded)
}
