package copperclient

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CreateDepositTarget create deposit target request
func (cc *CopperClient) CreateDepositTarget(currency string) (*DepositTarget, error) {
	isSupported := cfg.IsSupportedCurrency(currency)
	if !isSupported {
		return nil, ErrUnknownCurrency
	}

	reqBody := CreateDepositTargetRequest{
		ExternalId:   uuid.NewString(),
		PortfolioId:  cc.portfolioID,
		Currency:     currency,
		MainCurrency: cc.mainCurrency,
	}

	respBody := CreateDepositTargetResponse{}
	err := cc.request(http.MethodPost, PathDepositTarget, map[string]string{}, &reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	result := DepositTarget{
		DepositTargetId: respBody.DepositTargetId,
		ExternalId:      respBody.ExternalId,
		Address:         *respBody.Address,
		Status:          respBody.Status,
		Currency:        currency,
	}

	if respBody.Currency == nil {
		result.Currency = respBody.MainCurrency
	} else {
		result.Currency = *respBody.Currency
	}

	return &result, nil
}

// GetCurrency get currency information
func (cc *CopperClient) GetCurrency(currency string) (*Currency, error) {
	isSupported := cfg.IsSupportedCurrency(currency)
	if !isSupported {
		return nil, ErrUnknownCurrency
	}

	respBody := CurrenciesResponse{}
	err := cc.request(
		http.MethodGet,
		PathCurrency,
		map[string]string{"currency": currency, "rateFor": cfg.RateCurrency()},
		nil,
		&respBody,
	)
	if err != nil {
		return nil, err
	}

	return &respBody.Currencies[0], nil
}

// CreateWithdrawOrder create withdraw order
func (cc *CopperClient) CreateWithdrawOrder(externalOrderID, toAddress, currency string, amount decimal.Decimal) (*Order, error) {
	isSupported := cfg.IsSupportedCurrency(currency)
	if !isSupported {
		return nil, ErrUnknownCurrency
	}

	reqBody := CreateOrderRequest{
		ExternalOrderId: externalOrderID,
		OrderType:       OrderTypeWithdraw,
		BaseCurrency:    currency,
		MainCurrency:    cfg.MainCurrency(),
		Amount:          amount.String(),
		PortfolioId:     cfg.PortfolioID(),
		ToAddress:       toAddress,
		Description:     externalOrderID,
	}

	respBody := CreateOrderResponse{}
	err := cc.request(http.MethodPost, PathOrder, map[string]string{}, &reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	convAmount, err := decimal.NewFromString(respBody.Amount)
	if err != nil {
		return nil, err
	}

	convFee, err := decimal.NewFromString(respBody.Extra.FeesPercent)
	if err != nil {
		return nil, err
	}

	return &Order{
		OrderId:         respBody.OrderId,
		ExternalOrderId: respBody.ExternalOrderId,
		Status:          respBody.Status,
		OrderType:       respBody.OrderType,
		Amount:          convAmount,
		Currency:        respBody.BaseCurrency,
		ToAddress:       respBody.Extra.ToAddress,
		FeesPercent:     convFee,
	}, nil
}

// CancelOrder cancel order
func (cc *CopperClient) CancelOrder(orderID string) error {
	return cc.request(http.MethodPatch, fmt.Sprintf(PathCancelOrder, orderID), map[string]string{}, nil, nil)
}

// GetOrdersList get orders list with limit and offset
func (cc *CopperClient) GetOrdersList(limit, offset int) ([]*OrderListItem, error) {
	respBody := OrdersListResponse{}
	err := cc.request(
		http.MethodGet,
		PathOrder,
		map[string]string{"limit": strconv.Itoa(limit), "offset": strconv.Itoa(offset)},
		nil,
		&respBody,
	)
	if err != nil {
		return nil, err
	}

	return respBody.Orders, nil
}
