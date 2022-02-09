package copperclient

import "time"

// TODO: внести в конфиг
const (
	DefClientTimeout = 10 * time.Second
	DefSecretKey     = ""
	DefAPIKey        = ""
	DefBaseURL       = "https://api.testnet.copper.co"
	DefPortfolioID   = ""
	DefMainCurrency  = CurrencySOL
	DefRateCurrency  = CurrencyUSD
)

// headers for copper requests
const (
	HeaderKeyAuthorization    = "Authorization"
	HeaderPrefixAuthorization = "ApiKey"
	HeaderKeyTimestamp        = "X-Timestamp"
	HeaderKeySignature        = "X-Signature"
	HeaderKeyContentType      = "Content-Type"
	HeaderValueContentType    = "application/json"
)

// copper API paths
const (
	PathOrder         = "/platform/orders"
	PathCancelOrder   = "/platform/orders/%s"
	PathDepositTarget = "/platform/deposit-targets"
	PathCurrency      = "/platform/currencies"
)

const (
	CurrencySOL = "SOL"
	CurrencyUSD = "USD"

	OrderTypeWithdraw = "withdraw"
	OrderTypeDeposit  = "deposit"
)

var supportedCurrencies = map[string]bool{
	CurrencySOL: true,
}

var cfg Config

type Config struct {
}

func (c Config) ClientTimeout() time.Duration {
	return DefClientTimeout
}

func (c Config) SecretKey() string {
	return DefSecretKey
}

func (c Config) APIKey() string {
	return DefAPIKey
}

func (c Config) BaseURL() string {
	return DefBaseURL
}

func (c Config) PortfolioID() string {
	return DefPortfolioID
}

func (c Config) MainCurrency() string {
	return DefMainCurrency
}

func (c Config) IsSupportedCurrency(currency string) bool {
	isSupported, ok := supportedCurrencies[currency]
	return ok && isSupported
}

func (c Config) RateCurrency() string {
	return DefRateCurrency
}
