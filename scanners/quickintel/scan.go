package quickintel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// QuickIntelResponse represents the structure of the response from QuickIntel API.
type QuickIntelResponse struct {
	TokenDetails struct {
		TokenName        string `json:"tokenName"`
		TokenSymbol      string `json:"tokenSymbol"`
		TokenDecimals    int    `json:"tokenDecimals"`
		TokenLogo        string `json:"tokenLogo"`
		TokenOwner       string `json:"tokenOwner"`
		TokenSupply      int    `json:"tokenSupply"`
		TokenCreatedDate int64  `json:"tokenCreatedDate"`
		QuickiTokenHash  struct {
			ExactQHash   string `json:"exact_qHash"`
			SimilarQHash string `json:"similar_qHash"`
		} `json:"quickiTokenHash"`
	} `json:"tokenDetails"`
	TokenDynamicDetails struct {
		LastUpdatedTimestamp int64 `json:"lastUpdatedTimestamp"`
		IsHoneypot           bool  `json:"is_Honeypot"`
	} `json:"tokenDynamicDetails"`
	QuickiAudit struct {
		ContractCreator           string   `json:"contract_Creator"`
		ContractOwner             string   `json:"contract_Owner"`
		ContractName              string   `json:"contract_Name"`
		ContractChain             string   `json:"contract_Chain"`
		ContractAddress           string   `json:"contract_Address"`
		ContractRenounced         bool     `json:"contract_Renounced"`
		IsLaunchpadContract       bool     `json:"is_Launchpad_Contract"`
		LaunchpadDetails          string   `json:"launchpad_Details"`
		HiddenOwner               bool     `json:"hidden_Owner"`
		HiddenOwnerModifiers      string   `json:"hidden_Owner_Modifiers"`
		IsProxy                   bool     `json:"is_Proxy"`
		ProxyImplementation       string   `json:"proxy_Implementation"`
		HasExternalContractRisk   bool     `json:"has_External_Contract_Risk"`
		ExternalContracts         string   `json:"external_Contracts"`
		HasObfuscatedAddressRisk  bool     `json:"has_Obfuscated_Address_Risk"`
		ObfuscatedAddressList     string   `json:"obfuscated_Address_List"`
		CanMint                   bool     `json:"can_Mint"`
		CantMintRenounced         string   `json:"cant_Mint_Renounced"`
		CanBurn                   bool     `json:"can_Burn"`
		CanBlacklist              bool     `json:"can_Blacklist"`
		CantBlacklistRenounced    bool     `json:"cant_Blacklist_Renounced"`
		CanMultiBlacklist         bool     `json:"can_MultiBlacklist"`
		CanWhitelist              bool     `json:"can_Whitelist"`
		CantWhitelistRenounced    bool     `json:"cant_Whitelist_Renounced"`
		CanUpdateFees             bool     `json:"can_Update_Fees"`
		CantUpdateFeesRenounced   bool     `json:"cant_Update_Fees_Renounced"`
		CanUpdateMaxWallet        bool     `json:"can_Update_Max_Wallet"`
		CantUpdateMaxWalletRen    bool     `json:"cant_Update_Max_Wallet_Renounced"`
		CanUpdateMaxTx            bool     `json:"can_Update_Max_Tx"`
		CantUpdateMaxTxRen        bool     `json:"cant_Update_Max_Tx_Renounced"`
		CanPauseTrading           bool     `json:"can_Pause_Trading"`
		CantPauseTradingRen       bool     `json:"cant_Pause_Trading_Renounced"`
		HasTradingCooldown        bool     `json:"has_Trading_Cooldown"`
		CanUpdateWallets          bool     `json:"can_Update_Wallets"`
		HasSuspiciousFunctions    bool     `json:"has_Suspicious_Functions"`
		HasExternalFunctions      bool     `json:"has_External_Functions"`
		HasFeeWarning             bool     `json:"has_Fee_Warning"`
		HasModifiedTransferWarn   bool     `json:"has_ModifiedTransfer_Warning"`
		ModifiedTransferFuncs     string   `json:"modified_Transfer_Functions"`
		SuspiciousFuncs           string   `json:"suspicious_Functions"`
		ExternalFuncs             []string `json:"external_Functions"`
		FeeUpdateFuncs            []string `json:"fee_Update_Functions"`
		HasScams                  bool     `json:"has_Scams"`
		MatchedScams              string   `json:"matched_Scams"`
		ScamFuncs                 string   `json:"scam_Functions"`
		HasKnownScamWalletFunding bool     `json:"has_Known_Scam_Wallet_Funding"`
		KnownScamWalletFunding    string   `json:"known_Scam_Wallet_Funding"`
		ContractLinks             []string `json:"contract_Links"`
		Functions                 []string `json:"functions"`
		OnlyOwnerFunctions        []string `json:"onlyOwner_Functions"`
		MultiBlacklistFuncs       string   `json:"multiBlacklistFunctions"`
		HasGeneralVulnerabilities bool     `json:"has_General_Vulnerabilities"`
		GeneralVulnerabilities    string   `json:"general_Vulnerabilities"`
	} `json:"quickiAudit"`
	ProjectVerified  bool        `json:"projectVerified"`
	KycVerifications interface{} `json:"kycVerifications"`
	ExternalAudits   interface{} `json:"externalAudits"`
}

// Scan sends a request to QuickIntel API to get information about a token
// identified by its hash. It returns the response received or an error if any.
func Scan(tokenHash string) (QuickIntelResponse, error) {
	var response QuickIntelResponse

	// URL and request method
	url := "https://app.quickintel.io/api/quicki/getquickiauditfull"
	method := "POST"

	// Prepare the request body
	request := fmt.Sprintf("{\"chain\":\"eth\",\"tokenAddress\":\"%v\",\"tier\":\"basic\"}", tokenHash)
	payload := strings.NewReader(request)

	// Create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return response, err
	}
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	// Unmarshal JSON response into QuickIntelResponse struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
