package utils

import (
	"bufio"
	"os"
	"strings"
)

func RemoveHexPrefix(hexString string) string {
	// Check if the hex string has the "0x" prefix
	if strings.HasPrefix(hexString, "0x") || strings.HasPrefix(hexString, "0X") {
		// Remove the prefix and return the rest of the string
		return hexString[2:]
	}
	// If there is no "0x" prefix, return the original string
	return hexString
}

func ReplaceEnvValue(filename, key, newValue string) error {
	// Open the .env file for reading
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a temporary file for writing
	tempFile, err := os.CreateTemp("", "tempenv")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	// Create a scanner to read the .env file
	scanner := bufio.NewScanner(file)

	// Iterate through each line in the .env file
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			// Check if the current line contains the key we want to replace
			if parts[0] == key {
				// Replace the old value with the new value
				line = key + "=" + newValue
			}
		}

		// Write the line to the temporary file
		_, _ = tempFile.WriteString(line + "\n")
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return err
	}

	// Rename the temporary file to the original filename
	err = os.Rename(tempFile.Name(), filename)
	if err != nil {
		return err
	}

	return nil
}
