package token

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/models"
	"github.com/s-Amine/token-scan/scanners/ishoneypot"
	"github.com/s-Amine/token-scan/scanners/quickintel"
)

// TokenInfo represents information about a token.
type TokenInfo struct {
	TokenName                  string `json:"token_name,omitempty"`
	TokenSymbol                string `json:"token_symbol,omitempty"`
	Decimals                   int    `json:"decimals,omitempty"`
	UniswapV2Pair              string `json:"uniswapv2_pair,omitempty"`
	IsHoneypot                 bool   `json:"is_honeypot,omitempty"`
	IsOpenSource               bool   `json:"is_open_source,omitempty"`
	IsWhitelisted              bool   `json:"is_whitelisted,omitempty"`
	CanTakeBackOwnership       bool   `json:"can_take_back_ownership,omitempty"`
	OwnerChangeBalance         bool   `json:"owner_change_balance,omitempty"`
	CannotBuy                  bool   `json:"cannot_buy,omitempty"`
	CannotSellAll              bool   `json:"cannot_sell_all,omitempty"`
	IsMintable                 bool   `json:"is_mintable,omitempty"`
	HiddenOwner                bool   `json:"hidden_owner,omitempty"`
	TransferPausable           bool   `json:"transfer_pausable,omitempty"`
	IsBlacklisted              bool   `json:"is_blacklisted,omitempty"`
	BuyTax                     string `json:"buy_tax,omitempty"`
	SellTax                    string `json:"sell_tax,omitempty"`
	ExternalCall               bool   `json:"external_call,omitempty"`
	TradingCooldown            bool   `json:"trading_cooldown,omitempty"`
	PersonalSlippageModifiable bool   `json:"personal_slippage_modifiable,omitempty"`
}

// ToJSON converts TokenInfo to JSON string.
func (t *TokenInfo) ToJSON() (string, error) {
	jsonData, err := json.Marshal(t)
	if err != nil {
		return "", fmt.Errorf("error marshaling TokenInfo to JSON: %v", err)
	}
	return string(jsonData), nil
}

// setBoolField sets a boolean field based on the given value.
func (t *TokenInfo) setBoolField(value string, field *bool) {
	if value == "" {
		*field = false
	} else {
		val, err := strconv.ParseBool(value)
		if err != nil {
			fmt.Printf("Error parsing boolean value: %v\n", err)
			*field = false
		} else {
			*field = val
		}
	}
}

// InitTokenInfoFromQuickIntelResponse initializes TokenInfo from QuickIntelResponse.
func InitTokenInfoFromQuickIntelResponse(response quickintel.QuickIntelResponse) *TokenInfo {
	tokenInfo := &TokenInfo{
		TokenName:                  response.TokenDetails.TokenName,
		TokenSymbol:                response.TokenDetails.TokenSymbol,
		Decimals:                   response.TokenDetails.TokenDecimals,
		IsHoneypot:                 response.QuickiAudit.HiddenOwner,
		IsWhitelisted:              response.QuickiAudit.CanWhitelist,
		IsMintable:                 response.QuickiAudit.CanMint,
		TransferPausable:           response.QuickiAudit.CanPauseTrading,
		IsBlacklisted:              response.QuickiAudit.CanBlacklist,
		ExternalCall:               response.QuickiAudit.HasExternalContractRisk,
		TradingCooldown:            response.QuickiAudit.HasTradingCooldown,
		PersonalSlippageModifiable: false,
	}

	return tokenInfo
}

// InitTokenInfoFromGoPlus initializes TokenInfo from GoPlus response.
func InitTokenInfoFromGoPlus(r models.ResponseWrapperTokenSecurityResultAnon) *TokenInfo {
	tokenInfo := &TokenInfo{
		TokenName:   r.TokenName,
		TokenSymbol: r.TokenSymbol,
		BuyTax:      r.BuyTax,
		SellTax:     r.SellTax,
	}

	tokenInfo.setBoolField(r.CanTakeBackOwnership, &tokenInfo.CanTakeBackOwnership)
	tokenInfo.setBoolField(r.CannotBuy, &tokenInfo.CannotBuy)
	tokenInfo.setBoolField(r.CannotSellAll, &tokenInfo.CannotSellAll)
	tokenInfo.setBoolField(r.ExternalCall, &tokenInfo.ExternalCall)
	tokenInfo.setBoolField(r.HiddenOwner, &tokenInfo.HiddenOwner)
	tokenInfo.setBoolField(r.IsBlacklisted, &tokenInfo.IsBlacklisted)
	tokenInfo.setBoolField(r.IsHoneypot, &tokenInfo.IsHoneypot)
	tokenInfo.setBoolField(r.IsMintable, &tokenInfo.IsMintable)
	tokenInfo.setBoolField(r.IsOpenSource, &tokenInfo.IsOpenSource)
	tokenInfo.setBoolField(r.IsWhitelisted, &tokenInfo.IsWhitelisted)
	tokenInfo.setBoolField(r.OwnerChangeBalance, &tokenInfo.OwnerChangeBalance)
	tokenInfo.setBoolField(r.TradingCooldown, &tokenInfo.TradingCooldown)
	tokenInfo.setBoolField(r.TransferPausable, &tokenInfo.TransferPausable)

	return tokenInfo
}

// InitTokenInfoFromHoneypotResponse initializes TokenInfo from Honeypot response.
func InitTokenInfoFromHoneypotResponse(response ishoneypot.HoneypotResponse) *TokenInfo {
	tokenInfo := &TokenInfo{
		TokenName:     response.Token.Name,
		TokenSymbol:   response.Token.Symbol,
		Decimals:      response.Token.Decimals,
		UniswapV2Pair: response.Pair.PairAddress,
		IsHoneypot:    response.HoneypotResult.IsHoneypot,
		IsOpenSource:  response.ContractCode.OpenSource,
	}

	return tokenInfo
}

// UnifyTokenInfo unifies multiple TokenInfo instances into one.
func UnifyTokenInfo(info1, info2, info3 *TokenInfo) *TokenInfo {
	unifiedInfo := &TokenInfo{}

	// Helper function to determine worst-case boolean value
	worstBool := func(a, b, c bool) bool {
		return a || b || c
	}

	// Helper function to determine worst-case string value
	worstString := func(a, b, c string) string {
		max := a
		if len(b) > len(max) {
			max = b
		}
		if len(c) > len(max) {
			max = c
		}
		return max
	}

	// Helper function to determine worst-case integer value
	worstInt := func(a, b, c int) int {
		max := a
		if b > max {
			max = b
		}
		if c > max {
			max = c
		}
		return max
	}

	// Set values based on worst-case scenario
	unifiedInfo.TokenName = worstString(info1.TokenName, info2.TokenName, info3.TokenName)
	unifiedInfo.TokenSymbol = worstString(info1.TokenSymbol, info2.TokenSymbol, info3.TokenSymbol)
	unifiedInfo.Decimals = worstInt(info1.Decimals, info2.Decimals, info3.Decimals)
	unifiedInfo.UniswapV2Pair = worstString(info1.UniswapV2Pair, info2.UniswapV2Pair, info3.UniswapV2Pair)
	unifiedInfo.IsHoneypot = worstBool(info1.IsHoneypot, info2.IsHoneypot, info3.IsHoneypot)
	unifiedInfo.IsOpenSource = worstBool(info1.IsOpenSource, info2.IsOpenSource, info3.IsOpenSource)
	unifiedInfo.IsWhitelisted = worstBool(info1.IsWhitelisted, info2.IsWhitelisted, info3.IsWhitelisted)
	unifiedInfo.CanTakeBackOwnership = worstBool(info1.CanTakeBackOwnership, info2.CanTakeBackOwnership, info3.CanTakeBackOwnership)
	unifiedInfo.OwnerChangeBalance = worstBool(info1.OwnerChangeBalance, info2.OwnerChangeBalance, info3.OwnerChangeBalance)
	unifiedInfo.CannotBuy = worstBool(info1.CannotBuy, info2.CannotBuy, info3.CannotBuy)
	unifiedInfo.CannotSellAll = worstBool(info1.CannotSellAll, info2.CannotSellAll, info3.CannotSellAll)
	unifiedInfo.IsMintable = worstBool(info1.IsMintable, info2.IsMintable, info3.IsMintable)
	unifiedInfo.HiddenOwner = worstBool(info1.HiddenOwner, info2.HiddenOwner, info3.HiddenOwner)
	unifiedInfo.TransferPausable = worstBool(info1.TransferPausable, info2.TransferPausable, info3.TransferPausable)
	unifiedInfo.IsBlacklisted = worstBool(info1.IsBlacklisted, info2.IsBlacklisted, info3.IsBlacklisted)
	unifiedInfo.BuyTax = worstString(info1.BuyTax, info2.BuyTax, info3.BuyTax)
	unifiedInfo.SellTax = worstString(info1.SellTax, info2.SellTax, info3.SellTax)
	unifiedInfo.ExternalCall = worstBool(info1.ExternalCall, info2.ExternalCall, info3.ExternalCall)
	unifiedInfo.TradingCooldown = worstBool(info1.TradingCooldown, info2.TradingCooldown, info3.TradingCooldown)
	unifiedInfo.PersonalSlippageModifiable = false

	return unifiedInfo
}
