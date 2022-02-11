package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	//acquired by $ go get github.com/labstack/echo/v4
	"github.com/labstack/echo/v4"
)

const (
	// Config
	frontend_endpoint = "https://stage.bitgin.app/fiat-as-a-service"
	backend_endpoint  = "https://api.bitgin.app"
	key               = "<API_KEY>"    // API_KEY acquired from BITGIN
	secret            = "<SECRET_KEY>" // SECRET_KEY acquired from BITGIN
)

type RequestBodyPay struct {
	OrderID  *string `json:"order_id,omitempty"`
	Amount   *string `json:"amount,omitempty"`
	Address  string  `json:"address"`
	Chain    string  `json:"chain"`
	Currency string  `json:"currency"`
}

type ResponseBodyPay struct {
	URL string `json:"url"`
}

type RequestBodyReceipt struct {
	Pagination pagination `json:"pagination"`

	OrderID       *string  `json:"order_id"`
	Currency      []string `json:"currency"`
	StartDateUnix *int64   `json:"start_date"`
	EndDateUnix   *int64   `json:"end_date"`
}
type pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type faasPaymentReceipt struct {
	PaymentID string    `json:"payment_id"`
	OrderID   *string   `json:"order_id,omitempty"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Withdrawal faasWithdrawal `json:"withdrawal"`
}

type faasWithdrawal struct {
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

type apiResponseReceipt struct {
	Success   bool                 `json:"success"`
	Message   *string              `json:"message,omitempty"`
	Data      []faasPaymentReceipt `json:"data,omitempty"`
	RequestID *string              `json:"request_id,omitempty"`
}

func sign(payload string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(payload))
	signature := hex.EncodeToString(hash.Sum(nil))
	return signature
}

func main() {
	e := echo.New()
	e.POST("/bitgin-pay-url", func(c echo.Context) error {

		body := RequestBodyPay{}
		if err := c.Bind(&body); err != nil {
			log.Println("error: ", err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}

		data, _ := json.Marshal(&body)

		method := strings.ToUpper(http.MethodPost)
		path := "/faas/v1/pay"
		nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)

		payload := nonce + method + path + string(data)
		fmt.Println("payload: ", payload)

		signature := sign(payload)
		fmt.Println("signature: ", signature)

		u, _ := url.Parse(frontend_endpoint)
		q := u.Query()
		q.Add("key", key)
		q.Add("sign", signature)
		q.Add("nonce", nonce)
		q.Add("body", base64.RawURLEncoding.EncodeToString(data))
		u.RawQuery = q.Encode()

		log.Println("url: ", u.String())

		res := ResponseBodyPay{
			URL: u.String(),
		}

		return c.JSON(http.StatusOK, res)
	})

	e.POST("/bitgin-receipt", func(c echo.Context) error {

		body := RequestBodyReceipt{}
		if err := c.Bind(&body); err != nil {
			log.Println("error: ", err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}
		data, _ := json.Marshal(body)

		method := strings.ToUpper(http.MethodPost)
		path := "/faas/v1/receipt"
		nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)

		payload := nonce + method + path + string(data)
		fmt.Println("payload: ", payload)

		signature := sign(payload)
		fmt.Println("signature: ", signature)

		log.Printf("key:%s&sign:%s&nonce:%s", key, signature, nonce)

		req, err := http.NewRequest("POST", backend_endpoint+path, bytes.NewBuffer(data))
		if err != nil {
			log.Println("1")
			return c.String(http.StatusBadRequest, err.Error())
		}

		reqHeader := map[string]string{
			"BG-API-KEY":   key,
			"BG-API-SIGN":  signature,
			"BG-API-NONCE": nonce,
			"Content-Type": "application/json; charset=UTF-8",
		}

		for k, v := range reqHeader {
			req.Header.Set(k, v)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		respbody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		var apiRes apiResponseReceipt
		err = json.Unmarshal(respbody, &apiRes)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(resp.StatusCode, apiRes)
	})

	e.Logger.Fatal(e.Start(":8881"))
}
