package model

import "time"

type RequestBodyPay struct {
	OrderID  *string  `json:"order_id,omitempty"`
	Amount   *float64 `json:"amount,omitempty"`
	Address  string   `json:"address"`
	Chain    string   `json:"chain"`
	Currency string   `json:"currency"`
}

type ResponseBodyPay struct {
	URL string `json:"url"`
}

type RequestBodyReceipt struct {
	Pagination

	OrderID       *string `query:"order_id"`
	Currency      *string `query:"currency"`
	StartDateUnix *int64  `query:"start_date"`
	EndDateUnix   *int64  `query:"end_date"`
}
type Pagination struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
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
	UserID  string  `json:"user_id"`
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

type ApiResponseMineShare struct {
	Success   bool    `json:"success"`
	Message   *string `json:"message,omitempty"`
	RequestID *string `json:"request_id,omitempty"`
}
