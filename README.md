# Table of contents
- [API Authentication](#api-authentication)

- [Fiat-as-a-Service](#fiat-as-a-service)
  - REST API
    - [Get Payment Embedded URL](#get-payment-embedded-url)
    - [Get Receipts](#get-receipts)

- [Mine Share Service](#mine-share-service)
  - REST API
    - [Query BITGIN Addresses](#query-bitgin-addresses)
    - [Mine Share](#mine-share)

- [Mock Server](#mock-server)
  - [How to compile](#how-to-compile)
  - [Setup configuration](#setup-configuration)
  - [REST API of Mock Server](#rest-api-of-mock-server)
	

# API Authentication
- The first things you need to use the FaaS is your API key and your secret key.
- You can use API key and secret key to generate the necessary parameters that contain  `SIGN` , `KEY` , `NONCE`, `TIMESTAMP`

### How to sign your data ?
- Please refer to the code snippet on the github project to know how to sign your data.
	- [Go](https://github.com/BITGIN/bitgin-api-docs/blob/main/handler/handler.go#L24)
  
```go
import (
  "bytes"
  "net/http"
  "strconv"
  "crypto/hmac"
  "crypto/sha256"
  "encoding/hex"
  "math/rand"
  "time"
  "fmt"
)
func sign(payload string) string {
	hash := hmac.New(sha256.New, []byte("YOUR_API_SECRET"))
	hash.Write([]byte(payload))
	signature := hex.EncodeToString(hash.Sum(nil))
	return signature
}

func randFunc() string {
	rand.Seed(time.Now().Unix())
	// 2^32
	x := rand.Int63n(4294967296)
	return fmt.Sprintf("%08x", x)
}

func main() {
  method := http.MethodPost
  path := "<path_url>"
  timestamp := strconv.FormatInt(time.Now().Unix(), 10)
  nonce := randFunc()
  payload := fmt.Sprintf("%s%s%s%s%s", method, path, nonce, timestamp, string(<request_body>))
  signature := sign(payload)
  req, _ := http.NewRequest("POST", "<endpoint_url>" + path, bytes.NewBuffer(<request_body>))
  req.Header.Set("BG-API-KEY", "YOUR_API_KEY")
  req.Header.Set("BG-API-SIGN", signature)
  req.Header.Set("BG-API-NONCE", nonce)
  req.Header.Set("BG-API-TIMESTAMP", timestamp)
  req.Header.Set("Content-Type","application/json")
}

```

##### [Back to top](#table-of-contents)
# Fiat-as-a-Service

- The quickest way to build conversion-optimized cryptocurrency payment forms, hosted on BITGIN.

- Fiat-as-a-Service creates a secure, BITGIN-hosted payment page that lets you collect payments quickly. It works across cryptocurrency and can help increase your conversion. Fiat-as-a-Service makes it easy to build a first-class payments experience.

### How to contact BITGIN ? 
- [Contact us](https://bitgin.freshdesk.com/support/tickets/new)
### Try it now
- Use [Mock Server](#mock-server) to test BITGIN Fiat-as-a-Service right away.
  - Step 1: Deploy the Mock Server
  - Step 2: Call [Get Payment Embedded URL](#get-payment-embedded-url) to get BITGIN Frontend URL
  - Step 3: Open the URL then complete payment
  - Step 4: Call [Get Receipts](#get-receipts) to acquire payment receipts

##### [Back to top](#table-of-contents)



# REST API

## Get Payment Embedded URL

##### Request

**POST** `MOCK_SERVER_DOMAIN`/v1/faas/pay

> NOTE: The API is a tool to help you generate the embedded url that only implements in the [mock server](#mock-server), so it does not allow you to send request to BITGIN Server.

##### Headers
| Key | Value | Note |
| :---  | :--- | :---        |
| Content-Type | application/json | required, JSON Type |


##### Request

```json
{
    "order_id":"00001_1", //optional
    "address":"TXHzvoDBPaG7YbSgb3zdoosJK4x4Kmf2J2", 
    "chain":"Tron",
    "currency":"USDT", 
    "amount":15.345 //optional
}
```
##### Body Fields

| Field    | Type   | Note | Description |
| :---     | :---   | :--- | :---        |
| order_id | string | optional | Specify customize id |
| address  | string | required | Specify the address you want to deposit |
| chain    | string | required (Tron, Ethereum, Bitcoin...) | Specify valid chain |
| currency | string | required (USDT, ETH, BTC...) | Specify the valid currency |
| amount   | float | optional, Greater than 0 | Specify the amount |


> NOTE: `amount` is optional as it can be provided by customer
> 
> NOTE: `chain` and `currency` are case-sensitive


##### Response Format

An example of a successful response:

```json
{
  "url": "https://stage.bitgin.app/fiat-as-a-service?body=eyJvcmRlcl9pZCI6IjAwMDAxXzEiLCJhbW91bnQiOjE1LjM0NSwiYWRkcmVzcyI6IlRYSHp2b0RCUGFHN1liU2diM3pkb29zSks0eDRLbWYySjIiLCJjaGFpbiI6IlRyb24iLCJjdXJyZW5jeSI6IlVTRFQifQ&key=6OOphiLielvNYLmXL8Pj&nonce=9657a548&sign=c65efa5d1c711d4227e59f56298a85829b696bc777bd45e6b49a347faa5b7467&timestamp=1646546374"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| [url](#url-parameters) | string |  URL is the payment site provided by BITGIN|

</br>

**URL** 

https://`BITGIN_DOMAIN`/fiat-as-a-service?sign=`SIGN`&key=`KEY`&nonce=`NONCE`&timestamp=`TIMESTAMP`&body=`BODY`

The URL includes the following parameters:

##### URL Parameters

| Field | Note | Description |
| :---  | :--- | :---        |
| SIGN | [How to sign your data ?](https://github.com/BITGIN/bitgin-api-docs/blob/main/handler/handler.go#L24)  | Which is the SHA256 HMAC of the following four strings, using your API secret, as a hex string: Request timestamp (same as above), HTTP method in uppercase (e.g. GET or POST), Request path, including leading slash and any URL parameters but not including the hostname|
| KEY  | acquired from BITGIN | The key string for BITGIN verification merchant |
| NONCE | random number [0, 2^32) | random number in the half-open interval [0,2^32) with hexadecimal system |
| TIMESTAMP | represented by seconds |  Unix time of current time, the number of `seconds` elapsed since January 1, 1970 UTC |
| [Body](#body-format) | base64 encoding as specified by RFC 4648 | Specify the base64 encoded payment information |
##### [Back to top](#table-of-contents)


<br />

## Get Receipts

Get payment receipts 

##### Request 

**GET** /v1/faas/receipt?order_id={order_id}&currency={currency}&start_date={start_date}&end_date={end_end}&limit={limit}&offset={offset}

##### Headers
| Key | Value | Note |
| :---  | :--- | :---        |
| Content-Type | application/json | required, JSON Type |
##### Parameter

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| order_id | string | optional | Specify customize id |
| currency | string | optional (USDT, ETH, BTC...) | Specify the valid currency |
| start_date | int64 | optional | Specify the start date of query time interval |
| end_date | int64 | optional | Specify the end date of query time interval |
| limit | int | optional, Greater than or equal 0 | Specify the query limit |
| offset | int | optional, Greater than or equal 0 | Specify the query offset |

> NOTE: The query max limt is 500.
> 
> NOTE: `start_date` and `end_date` are Unix time, the number of seconds elapsed since January 1, 1970 UTC.
>
> NOTE: `currency` is case-sensitive

<br />

##### Response Format

An example of a successful response:
	
```json
{
  "success": true,
  "data" : [
        {
            "payment_id": "ad122e63-9112-499e-be60-1997f9455f6b",
            "order_id": "00001_1",
            "user_id": "19ffefc5-d286-448c-8e61-2946f61182e5",
            "created_at": "2022-02-10T11:59:26.33765Z",
            "updated_at": "2022-02-10T11:59:35.921309Z",
            "withdrawal": {
                "status": "completed",
                "completed_at": "2022-02-10T12:05:07.590928Z",
                "amount": "12.345678",
                "fee": "1",
                "currency": "USDT",
                "fee_currency": "USDT",
                "address": "TWJpcWeF3WQyp25hGwUvdB89wxjfFUmJgW",
                "chain": "Tron",
                "tx_id": "ec7073b61f7b653ff204d1ac916249b822824c3e0128df38b15cd10fe00b235e",
                "is_deduction": true
            }
        },
        {
            //.....
        }
  ]
}
```
	
An example of a fail response:

	
```json
{
  "success": false,
  "message": "invalid_request_format",
  "request_id": "NP3RiuzpV6gytwiArOEeFX2C7ao745rJ"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| success | bool | Success or not |
| message | string | Error message |
| [data](#data-format) | JSON | Receipts data |
| request_id | string | Request id For Debug  |


##### Data Format

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| payment_id | string |  ID of payment using by BITGIN|
| order_id | string |  ID of order (customize)|
| user_id | string |  ID of BIGIN user|
| created_at | string |  CreatedTime of payment|
| updated_at | string |  UpdatedTime of payment|
| status | string | Status of payment (completed, sent, pending)|
| completed_at | string |  CompletedTime of payment|
| amount | string |  Amount of payment|
| fee | string |  Fee of paymemt|
| currency | string | Currency of payment (USDT, ETH, BTC...)|
| fee_currency | string | Currency of fee of payment (USDT, ETH, BTC...)|
| address | string |  Deposit address of payment|
| chain | string |  Chain of payment (Tron, Ethereum, Bitcoin...)|
| tx_id | string | Transaction ID is a string that identifies a specific transaction on the blockchain|
| is_deduction | boolean | |


##### Error Code

| HTTP Code | Error | Description |
| :---      | :---  | :---        |
| 400  | invalid_api_request| Invalid request headers |
| 400  | invalid_request_format| Invalid request body |
| 401  | api_key_user_id_not_exist| API KEY not exist |
| 403  | permission_denied | No permission to call the API |
| 500  | unknown_error| Unkown error, please [Contact Us](https://bitgin.freshdesk.com/support/tickets/new) |
| 500  | unexpected_error| Unexpected error, please [Contact Us](https://bitgin.freshdesk.com/support/tickets/new) |



##### [Back to top](#table-of-contents)

<br />

# Mine Share Service

### How to contact BITGIN ? 
- [Contact us](https://bitgin.freshdesk.com/support/tickets/new)
### Try it now

- Use [Mock Server](#mock-server) to test BITGIN Mine Share Service right away.
  - Step 1: Deploy the Mock Server
  - Step 2: Call [Query BITGIN Addresses API](#query-bitgin-addresses) to get BITGIN addresses with user_id
  - Step 3: Deposit total amount to the designated exclusive address and get the txid
  - Step 4: Call [Mine Share API](#mine-share) to distribute the total amount to addresses that you specify
> NOTE: After confirmed that the transaction in Step 3 is completed, you need to wait a few minutes (recommend 10 miniutes) to enter Step 4 until BITGIN receive your deposit Information from our pool address, otherwise you will receive `txid_not_found` error
>
##### [Back to top](#table-of-contents)


# REST API
## Query BITGIN Addresses


##### Request

**POST** /v1/mine/query

##### Headers
| Key | Value | Note |
| :---  | :--- | :---        |
| Content-Type | application/json | required, JSON Type |
###### Post body


```json

{
  "currency": "ETH",
  "addresses": [
    "0x14545e3C46aDf35673E2483c3EE957bdb5aF7311",
    "0xE8e6ee727a74488631448e3624A0D80B11A431B0",
    //...
  ]
}

```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| currency | string | required (ETH, TRX, BTC...) | Specify currency type |
| addresses | string array | required | Specify addresses you want to check are BITGIN addresses or not |

##### Response Format

An example of a successful response:
	
```json
{
    "success": true,
    "data": {
        "bitgin_addresses": [
            {
                "user_id": "351c0599-17b9-44ad-b10e-29f93b52863e",
                "address": "0x14545e3C46aDf35673E2483c3EE957bdb5aF7311"
            },
            {
                "user_id": "1f69e74f-f296-458b-bcfb-8e5ee3232969",
                "address": "0xE8e6ee727a74488631448e3624A0D80B11A431B0"
            }
        ]
    }
}
```
	
An example of a fail response:

	
```json
{
  "success": false,
  "message": "invalid_request_format",
  "request_id": "NP3RiuzpV6gytwiArOEeFX2C7ao745rJ"
}
```
The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| success | bool | Success or not |
| message | string | Error message |
| [data](#data-format-for-query-bitgin-addresses) | JSON | Receipts data |
| request_id | string | Request id For Debug  |

##### Data Format For Query BITGIN Addresses

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| user_id | string |  ID of BIGIN user|
| address | string |  Address of BITGIN user|

##### [Back to top](#table-of-contents)
<br />

## Mine Share


##### Request

**POST** /v1/mine/share

##### Headers
| Key | Value | Note |
| :---  | :--- | :---        |
| Content-Type | application/json | required, JSON Type |

###### Post body


```json

{
  "txid": "0xe3e06dfefd94e7ea3b267445505369695531ce00c4c14b165d0d7c4b586dc181",
  "share": [
    {
      "user_id": "351c0599-17b9-44ad-b10e-29f93b52863e",
      "address": "0x14545e3C46aDf35673E2483c3EE957bdb5aF7311",
      "amount": 2.345
    },
    {
      "user_id": "1f69e74f-f296-458b-bcfb-8e5ee3232969",
      "address": "0xE8e6ee727a74488631448e3624A0D80B11A431B0",
      "amount": 1.56
    },
    //...
  ]
}

```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| txid | string | required | Transaction ID is a string that identifies a specific transaction on the blockchain |
| user_id | string | required, aquired from [Query API](#query-bitgin-addresses) | ID of BIGIN user|
| addresses | string array | required | Specify BITGIN address corresponding to user_id|
| amount | float | required |Amount of mine share for the address|

> NOTE: The sum of amount must be the same as the txid amount.
>

##### Response Format

An example of a successful response:
	
```json
{
    "success": true,
}
```
	
An example of a fail response:

	
```json
{
  "success": false,
  "message": "invalid_request_format",
  "request_id": "NP3RiuzpV6gytwiArOEeFX2C7ao745rJ"
}
```
The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| success | bool | Success or not |
| message | string | Error message |
| request_id | string | Request id For Debug  |

##### [Back to top](#table-of-contents)


<br />

# Mock Server

## How to compile

- Clone project
```
$ git clone https://github.com/BITGIN/faas-api-docs.git
```

- Install [Go](https://go.dev/dl/)

- [Setsup Configuration](#setup-configuration)

- Run mock server
```
$ go run main.go
```

## Setup configuration

>	NOTE: Configure in /config/config.go
```go
Frontend_Endpoint = "<FRONTEND_ENDPOINT>"
Backend_Endpoint  = "<BACKEND_ENDPOINT>"
Key               = "<API_KEY>"  // API_KEY acquired from BITGIN
Secret            = "<SECRET_KEY>" // SECRET_KEY acquired from BITGIN
```

## REST API of Mock Server

- ## Fiat-as-a-Service
### Get Payment Embedded URL
```
http://localhost:8888/v1/faas/pay
```

[API definition](#get-payment-embedded-url)


### Get Receipts

```
http://localhost:8888/v1/faas/receipt
```

[API definition](#get-receipts)

<br />

- ## Mine Share Service

### Query BITGIN Addresses

```
http://localhost:8888/v1/mine/query
```

[API definition](#query-bitgin-addresses)

### Mine Share
```
http://localhost:8888/v1/mine/share
```
[API definition](#mine-share)

##### [Back to top](#table-of-contents)
