package copperclient

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

// CopperClientInterface copper client interface
type CopperClientInterface interface {
	CreateDepositTarget(currency string) (*DepositTarget, error)
	GetCurrency(currency string) *Currency
	CreateWithdrawOrder(orderID, toAddress, currency string, amount decimal.Decimal) (*Order, error)
	CancelOrder(orderID string) error
}

// ErrorResponse copper API error response
type ErrorResponse struct {
	Error   *string `json:"error,omitempty"`
	Message string  `json:"message"`
}

func (er *ErrorResponse) GetError() error {
	return errors.New(er.Message)
}

// Currency currency item for currencies response
type Currency struct {
	Currency      string   `json:"currency"`
	MainCurrency  string   `json:"mainCurrency"`
	Name          string   `json:"name"`
	Fiat          bool     `json:"fiat"`
	Priority      string   `json:"priority"`
	Confirmations string   `json:"confirmations"`
	Decimal       string   `json:"decimal"`
	Tags          []string `json:"tags"`
	Color         string   `json:"color"`
	FeesLevels    []struct {
		FeeLevel string `json:"feeLevel"`
		Value    string `json:"value"`
	} `json:"feesLevels"`
	ExplorerUrl string `json:"explorerUrl"`
	Embedded    struct {
		Price struct {
			BaseCurrency  string `json:"baseCurrency"`
			QuoteCurrency string `json:"quoteCurrency"`
			Price         string `json:"price"`
		} `json:"price"`
	} `json:"_embedded"`
}

// CurrenciesResponse currencies response
type CurrenciesResponse struct {
	Currencies []Currency `json:"currencies"`
}

// CreateDepositTargetRequest create deposit target request structure
type CreateDepositTargetRequest struct {
	ExternalId   string  `json:"externalId"`
	PortfolioId  string  `json:"portfolioId"`
	Currency     string  `json:"currency"`
	MainCurrency string  `json:"mainCurrency"`
	Name         *string `json:"name,omitempty"`
}

// CreateDepositTargetResponse create deposit target response structure
type CreateDepositTargetResponse struct {
	DepositTargetId         string  `json:"depositTargetId"`
	ExternalId              string  `json:"externalId"`
	PortfolioId             string  `json:"portfolioId"`
	PortfolioType           string  `json:"portfolioType"`
	OrganizationId          string  `json:"organizationId"`
	TargetType              string  `json:"targetType"`
	Name                    *string `json:"name,omitempty"`
	Address                 *string `json:"address,omitempty"`
	Memo                    *string `json:"memo,omitempty"`
	Status                  string  `json:"status"`
	Currency                *string `json:"currency,omitempty"`
	MainCurrency            string  `json:"mainCurrency"`
	CreatedBy               string  `json:"createdBy"`
	UpdatedBy               string  `json:"updatedBy"`
	AcceptAllTokens         bool    `json:"acceptAllTokens"`
	RequireTokensActivation bool    `json:"requireTokensActivation"`
	ActivatedTokens         *[]struct {
		Currency string `json:"currency"`
		Status   string `json:"status"`
	} `json:"activatedTokens,omitempty"`
}

// DepositTarget deposit target
type DepositTarget struct {
	DepositTargetId string `json:"deposit_target_id"`
	ExternalId      string `json:"external_id"`
	Address         string `json:"address"`
	Status          string `json:"status"`
	Currency        string `json:"currency"`
}

// CreateOrderRequest create order request body
type CreateOrderRequest struct {
	ExternalOrderId string `json:"externalOrderId"`
	OrderType       string `json:"orderType"`
	BaseCurrency    string `json:"baseCurrency"`
	MainCurrency    string `json:"mainCurrency"`
	Amount          string `json:"amount"`
	PortfolioId     string `json:"portfolioId"`
	ToAddress       string `json:"toAddress"`
	Description     string `json:"description"`
}

// CreateOrderResponse create order response
type CreateOrderResponse struct {
	OrderId         string `json:"orderId"`
	ExternalOrderId string `json:"externalOrderId"`
	Status          string `json:"status"`
	OrderType       string `json:"orderType"`
	PortfolioId     string `json:"portfolioId"`
	PortfolioType   string `json:"portfolioType"`
	AccountId       string `json:"accountId"`
	Amount          string `json:"amount"`
	BaseCurrency    string `json:"baseCurrency"`
	MainCurrency    string `json:"mainCurrency"`
	Extra           struct {
		Description string `json:"description"`
		ToAddress   string `json:"toAddress"`
		FeesPercent string `json:"feesPercent"`
	} `json:"extra"`
	CreatedBy      string `json:"createdBy"`
	OrganizationId string `json:"organizationId"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

// Order order information
type Order struct {
	OrderId         string          `json:"orderId"`
	ExternalOrderId string          `json:"external_order_id"`
	Status          string          `json:"status"`
	OrderType       string          `json:"order_type"`
	Amount          decimal.Decimal `json:"amount"`
	Currency        string          `json:"currency"`
	ToAddress       string          `json:"to_address"`
	FeesPercent     decimal.Decimal `json:"fees_percent"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type OrderListItem struct {
	OrderId         string `json:"orderId"`
	ExternalOrderId string `json:"externalOrderId"`
	Status          string `json:"status"`
	OrderType       string `json:"orderType"`
	PortfolioId     string `json:"portfolioId"`
	PortfolioType   string `json:"portfolioType"`
	AccountId       string `json:"accountId"`
	Amount          string `json:"amount"`
	BaseCurrency    string `json:"baseCurrency"`
	MainCurrency    string `json:"mainCurrency"`
	Extra           struct {
		Confirmations           *string   `json:"confirmations,omitempty"`
		TransactionId           *string   `json:"transactionId,omitempty"`
		FromAddresses           *[]string `json:"fromAddresses,omitempty"`
		DepositTargetId         *string   `json:"depositTargetId,omitempty"`
		TransferFees            *string   `json:"transferFees,omitempty"`
		TransferFeesCurrency    *string   `json:"transferFeesCurrency,omitempty"`
		TransferTransactionId   *string   `json:"transferTransactionId,omitempty"`
		TransferDepositTargetId *string   `json:"transferDepositTargetId,omitempty"`
		ToAddress               *string   `json:"toAddress,omitempty"`
		FeesPercent             *string   `json:"feesPercent,omitempty"`
	} `json:"extra"`
	OrganizationId string  `json:"organizationId"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      string  `json:"updatedAt"`
	TerminatedAt   string  `json:"terminatedAt"`
	CreatedBy      *string `json:"createdBy,omitempty"`
}

type OrdersListResponse struct {
	Orders []*OrderListItem `json:"orders"`
}
