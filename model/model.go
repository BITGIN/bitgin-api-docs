package model

import "time"

type RequestBodyPay struct {
	OrderID  *string `json:"order_id,omitempty"`
	Amount   *int    `json:"amount,omitempty"`
	Address  string  `json:"address"`
	Chain    string  `json:"chain"`
	Currency string  `json:"currency"`
}

type ResponseBodyPay struct {
	URL string `json:"url"`
}

type RequestBodyReceipt struct {
	Pagination Pagination `json:"pagination"`

	OrderID       *string  `json:"order_id"`
	Currency      []string `json:"currency"`
	StartDateUnix *int64   `json:"start_date"`
	EndDateUnix   *int64   `json:"end_date"`
}
type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type FaasPaymentReceipt struct {
	PaymentID string    `json:"payment_id"`
	OrderID   *string   `json:"order_id,omitempty"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Withdrawal FaasWithdrawal `json:"withdrawal"`
}

type FaasWithdrawal struct {
	Status      string     `json:"status"`
	CompletedAt *time.Time `json:"completed_at"`
	Amount      string     `json:"amount"`
	Fee         string     `json:"fee"`
	Currency    string     `json:"currency"`
	FeeCurrency string     `json:"fee_currency"`
	Address     string     `json:"address,omitempty"`
	Chain       string     `json:"chain,omitempty"`
	TxID        *string    `json:"tx_id,omitempty"`
	IsDeduction bool       `json:"is_deduction,omitempty"`
}

type ApiResponseFaasReceipt struct {
	Success   bool                 `json:"success"`
	Message   *string              `json:"message,omitempty"`
	Data      []FaasPaymentReceipt `json:"data,omitempty"`
	RequestID *string              `json:"request_id,omitempty"`
}

type MineCheckBitginAddressRequest struct {
	Currency  string   `json:"currency"`
	Addresses []string `json:"addresses"`
}
type MineCheckBitginAddressResponse struct {
	BitginAddresses []MineBitginUserInfo `json:"bitgin_addresses"`
}
type MineBitginUserInfo struct {
	UserID  string `json:"user_id"`
	Address string `json:"address"`
}

type ApiResponseMineQuery struct {
	Success   bool                           `json:"success"`
	Message   *string                        `json:"message,omitempty"`
	Data      MineCheckBitginAddressResponse `json:"data,omitempty"`
	RequestID *string                        `json:"request_id,omitempty"`
}

type MineShareReq struct {
	TxID  string          `json:"txid"`
	Share []MineShareInfo `json:"share"`
}

type MineShareInfo struct {
	UserID  string `json:"user_id"`
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}

type ApiResponseMineShare struct {
	Success   bool    `json:"success"`
	Message   *string `json:"message,omitempty"`
	RequestID *string `json:"request_id,omitempty"`
}
