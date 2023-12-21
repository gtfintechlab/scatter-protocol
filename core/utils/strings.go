package utils

import "strings"

func RemoveHexPrefix(hexString string) string {
	// Check if the hex string has the "0x" prefix
	if strings.HasPrefix(hexString, "0x") || strings.HasPrefix(hexString, "0X") {
		// Remove the prefix and return the rest of the string
		return hexString[2:]
	}
	// If there is no "0x" prefix, return the original string
	return hexString
}
