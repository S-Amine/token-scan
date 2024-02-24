package goplus

import (
	"fmt"

	"github.com/GoPlusSecurity/goplus-sdk-go/api/token"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/errorcode"
	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/models"
)

// Scan performs a security scan on a token identified by its hash.
// It returns the security result wrapped in a response structure.
func Scan(tokenHash string) (models.ResponseWrapperTokenSecurityResultAnon, error) {
	// Create a new token security instance
	tokenSecurity := token.NewTokenSecurity(nil)
	// Specify the chain ID
	chainId := "1"
	// Prepare the list of contract addresses for scanning
	contractAddresses := []string{tokenHash}
	// Run the security scan
	data, err := tokenSecurity.Run(chainId, contractAddresses)
	// Handle any errors that occur during the scan
	if err != nil {
		// Return the error if it exists
		return models.ResponseWrapperTokenSecurityResultAnon{}, err
	}
	// Check the response code for success
	if data.Payload.Code != errorcode.SUCCESS {
		// If the response code indicates failure, return an error with the message
		return models.ResponseWrapperTokenSecurityResultAnon{}, fmt.Errorf("scan failed: %s", data.Payload.Message)
	}
	// Retrieve the security result for the specified token hash
	value := data.Payload.Result[tokenHash]

	return value, nil
}
