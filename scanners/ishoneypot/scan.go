package ishoneypot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HoneypotResponse represents the structure of the response from Honeypot API.
type HoneypotResponse struct {
	Token          TokenInfo  `json:"token"`
	WithToken      TokenInfo  `json:"withToken"`
	Simulation     Simulation `json:"simulationResult"`
	HoneypotResult struct {
		IsHoneypot bool `json:"isHoneypot"`
	} `json:"honeypotResult"`
	HolderAnalysis struct {
		Holders         string    `json:"holders"`
		Successful      string    `json:"successful"`
		Failed          string    `json:"failed"`
		Siphoned        string    `json:"siphoned"`
		AverageTax      int       `json:"averageTax"`
		AverageGas      float64   `json:"averageGas"`
		HighestTax      int       `json:"highestTax"`
		HighTaxWallets  string    `json:"highTaxWallets"`
		TaxDistribution []TaxInfo `json:"taxDistribution"`
		SnipersFailed   int       `json:"snipersFailed"`
		SnipersSuccess  int       `json:"snipersSuccess"`
	} `json:"holderAnalysis"`
	ContractCode struct {
		OpenSource     bool `json:"openSource"`
		RootOpenSource bool `json:"rootOpenSource"`
		IsProxy        bool `json:"isProxy"`
		HasProxyCalls  bool `json:"hasProxyCalls"`
	} `json:"contractCode"`
	Chain       ChainInfo `json:"chain"`
	Router      string    `json:"router"`
	Pair        PairInfo  `json:"pair"`
	PairAddress string    `json:"pairAddress"`
}

// TokenInfo represents token information.
type TokenInfo struct {
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Decimals     int    `json:"decimals"`
	Address      string `json:"address"`
	TotalHolders int    `json:"totalHolders"`
}

// Simulation represents simulation data.
type Simulation struct {
	BuyTax      int    `json:"buyTax"`
	SellTax     int    `json:"sellTax"`
	TransferTax int    `json:"transferTax"`
	BuyGas      string `json:"buyGas"`
	SellGas     string `json:"sellGas"`
}

// TaxInfo represents tax information.
type TaxInfo struct {
	Tax   int `json:"tax"`
	Count int `json:"count"`
}

// ChainInfo represents chain information.
type ChainInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Currency  string `json:"currency"`
}

// PairInfo represents pair information.
type PairInfo struct {
	PairName           string  `json:"name"`
	PairAddress        string  `json:"address"`
	Token0             string  `json:"token0"`
	Token1             string  `json:"token1"`
	Type               string  `json:"type"`
	ChainId            string  `json:"chainId"`
	Reserves0          string  `json:"reserves0"`
	Reserves1          string  `json:"reserves1"`
	Liquidity          float64 `json:"liquidity"`
	Router             string  `json:"router"`
	CreatedAtTimestamp string  `json:"createdAtTimestamp"`
	CreationTxHash     string  `json:"creationTxHash"`
}

// Scan sends a request to Honeypot API to check if a token is a honeypot.
// It returns the response received or an error if any.
func Scan(tokenHash string) (HoneypotResponse, error) {
	// Construct the URL
	url := fmt.Sprintf("https://api.honeypot.is/v2/IsHoneypot?address=%v&chainID=1", tokenHash)
	method := "GET"

	// Create an HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return HoneypotResponse{}, err
	}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return HoneypotResponse{}, err
	}
	defer res.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return HoneypotResponse{}, err
	}

	// Unmarshal JSON response into HoneypotResponse struct
	var response HoneypotResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return HoneypotResponse{}, err
	}

	return response, nil
}
