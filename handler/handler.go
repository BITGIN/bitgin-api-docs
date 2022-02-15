package handler

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

	"github.com/BITGIN/bitgin-api-docs/config"
	"github.com/BITGIN/bitgin-api-docs/model"
	"github.com/labstack/echo/v4"
)

func sign(payload string) string {
	hash := hmac.New(sha256.New, []byte(config.Secret))
	hash.Write([]byte(payload))
	signature := hex.EncodeToString(hash.Sum(nil))
	return signature
}

func FaasPayHandler(c echo.Context) error {
	body := model.RequestBodyPay{}
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

	u, _ := url.Parse(config.Frontend_Endpoint)
	q := u.Query()
	q.Add("key", config.Key)
	q.Add("sign", signature)
	q.Add("nonce", nonce)
	q.Add("body", base64.RawURLEncoding.EncodeToString(data))
	u.RawQuery = q.Encode()

	log.Println("url: ", u.String())

	res := model.ResponseBodyPay{
		URL: u.String(),
	}
	return c.JSON(http.StatusOK, res)
}

func FaasReceiptHandler(c echo.Context) error {

	body := model.RequestBodyReceipt{}
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

	log.Printf("key:%s&sign:%s&nonce:%s", config.Key, signature, nonce)

	req, err := http.NewRequest("POST", config.Backend_Endpoint+path, bytes.NewBuffer(data))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	reqHeader := map[string]string{
		"BG-API-KEY":   config.Key,
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

	var apiRes model.ApiResponseFaasReceipt
	err = json.Unmarshal(respbody, &apiRes)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(resp.StatusCode, apiRes)
}

func MineQueryAddressesHandler(c echo.Context) error {

	body := model.MineCheckBitginAddressRequest{}
	if err := c.Bind(&body); err != nil {
		log.Println("error: ", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}
	data, _ := json.Marshal(body)

	method := strings.ToUpper(http.MethodPost)
	path := "/mine/v1/query"
	nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)

	payload := nonce + method + path + string(data)
	fmt.Println("payload: ", payload)

	signature := sign(payload)
	fmt.Println("signature: ", signature)

	log.Printf("key:%s&sign:%s&nonce:%s", config.Key, signature, nonce)

	req, err := http.NewRequest("POST", config.Backend_Endpoint+path, bytes.NewBuffer(data))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	reqHeader := map[string]string{
		"BG-API-KEY":   config.Key,
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

	var apiRes model.ApiResponseMineQuery
	err = json.Unmarshal(respbody, &apiRes)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(resp.StatusCode, apiRes)
}

func MineShareHandler(c echo.Context) error {

	body := model.MineShareReq{}
	if err := c.Bind(&body); err != nil {
		log.Println("error: ", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}
	data, _ := json.Marshal(body)

	method := strings.ToUpper(http.MethodPost)
	path := "/mine/v1/share"
	nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)

	payload := nonce + method + path + string(data)
	fmt.Println("payload: ", payload)

	signature := sign(payload)
	fmt.Println("signature: ", signature)

	log.Printf("key:%s&sign:%s&nonce:%s", config.Key, signature, nonce)

	req, err := http.NewRequest("POST", config.Backend_Endpoint+path, bytes.NewBuffer(data))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	reqHeader := map[string]string{
		"BG-API-KEY":   config.Key,
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

	var apiRes model.ApiResponseMineShare
	err = json.Unmarshal(respbody, &apiRes)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(resp.StatusCode, apiRes)
}
