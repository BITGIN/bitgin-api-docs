# Table of contents

- [Get Started](#get-started)
- [API Authentication](#api-authentication)
- [Services](#services)
    - [FaaS Payment Through BITGIN Frontend](#faas-payment-through-bitgin-frontend)
    - [FaaS Get Receipt From BITGIN Backend](#faas-get-receipt-from-bitgin-backend)
- [Mock Server](#mock-server)

	


# Get Started


### BITGIN Fiat-as-a-Service

- The quickest way to build conversion-optimized cryptocurrency payment forms, hosted on BITGIN.

- Fiat-as-a-Service creates a secure, BITGIN-hosted payment page that lets you collect payments quickly. It works across cryptocurrency and can help increase your conversion. Fiat-as-a-Service makes it easy to build a first-class payments experience.

### How to contact BITGIN ? 
- support@bitgin.net
### Try it now
- Use [Mock Server](#mock-server) to test BITGIN Fiat-as-a-Service right away.

##### [Back to top](#table-of-contents)


# API Authentication
- The first things you need to use the FaaS is your API key and your secret key.
- You can use API key and secret key to generate the necessary parameters that contain  `SIGN` , `KEY` , `NONCE`

### How to sign your data ?
- Please refer to the code snippet on the github project to know how to sign your data.
	- [Go](https://github.com/BITGIN/faas-api-docs/blob/main/main.go#L103)

##### [Back to top](#table-of-contents)

<br />

# Services

<br />


## FaaS Payment Through BITGIN Frontend

**URL** 

https://`BITGIN_DOMAIN`/fiat-as-a-service?sign=`SIGN`&key=`KEY`&nonce=`NONCE`&body=`BODY`

<br />

The URL includes the following parameters:

##### URL Parameters

| Field | Note | Description |
| :---  | :--- | :---        |
| SIGN | acquired from BITGIN | Which is the SHA256 HMAC of the following four strings, using your API secret, as a hex string: Request timestamp (same as above), HTTP method in uppercase (e.g. GET or POST), Request path, including leading slash and any URL parameters but not including the hostname|
| KEY  | acquired from BITGIN | The key string for BITGIN verification merchant |
| NONCE | represented by milliseconds |  Unix time of current time, the number of `milliseconds` elapsed since January 1, 1970 UTC |
| [Body](#body-format) | base64 encoding as specified by RFC 4648 | Specify the base64 encoded payment information |



##### Body Format

```json
{
    "order_id":"00001_1", //optional
    "address":"TXHzvoDBPaG7YbSgb3zdoosJK4x4Kmf2J2", 
    "chain":"Tron",
    "currency":"USDT", 
    "amount":150 //optional
}
```
##### Body Fields

| Field    | Type   | Note | Description |
| :---     | :---   | :--- | :---        |
| order_id | string | optional | Specify customize id |
| address  | string | required | Specify the address you want to deposit |
| chain    | string | required (Tron, Ethereum, Bitcoin...) | Specify valid chain |
| currency | string | required (USDT, ETH, BTC...) | Specify the valid currency |
| amount   | string | optional, Greater than 0 | Specify the amount |


> NOTE: `amount` is optional as `amount` can be provided by customer
> 

##### Response Format

An example of a successful response:

```json
{
  "url": "https://bitgin.net/fiat-as-a-service?sign=<sign>&key=<key>&nonce=<nonce>&body=<body>"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| url | string |  URL is the payment site provided by BITGIN|


##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |

##### [Back to top](#table-of-contents)


<br />
<br />

## FaaS Get Receipt From BITGIN Backend

Get payment receipts 

##### Request Method

**POST** 

##### Request Format

An example of the request:

###### API

```
/faas/v1/receipts
```

###### Post body

```json
{
  "order_id": "00002_1",
  "currency": [
    "USDT",
    "ETH",
    //...
  ],
  "start_date": 1644398340, //unix
  "end_date": 1644571140, //unix
  "pagination": {
      "limit": 15,
      "offset": 3
  } 
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| order_id | string | optional | Specify customize id |
| currency | array | optional (USDT, ETH, BTC...) | Specify the valid currency |
| start_date | int64 | optional | Specify the start date of query time interval |
| end_date | int64 | optional | Specify the end date of query time interval |
| limit | int | optional, Greater than or equal 0 | Specify the query limit |
| offset | int | optional, Greater than or equal 0 | Specify the query offset |

> NOTE: The query max limt is 500.
> 
> NOTE: `start_date` and `end_date` are Unix time, the number of seconds elapsed since January 1, 1970 UTC.

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
| currency | string | Currency of payment (USDT, ETH, Bitcoin...)|
| fee_currency | string | Currency of fee of payment (USDT, ETH, Bitcoin...)|
| address | string |  Deposit address of payment|
| chain | string |  Chain of payment (Tron, Ethereum, Bitcoin...)|
| tx_id | string | Transaction ID is a string that identifies a specific transaction on the blockchain|
| is_deduction | boolean | |

##### [Back to top](#table-of-contents)

<br />
<br />

# Mock Server

### How to compile

- Clone project
```
$ git clone https://github.com/BITGIN/faas-api-docs.git
```

- Install [Go](https://go.dev/dl/)


- Run mock server
```
$ go run main.go
```

### Setup configuration

>	NOTE: Configure in main.go
```go
key    = "<API_KEY>"  // API_KEY acquired from BITGIN
secret = "<SECRET_KEY>" // SECRET_KEY acquired from BITGIN
```

### Services

- FaaS Payment Through BITGIN Frontend
##### POST Method
```
/bitgin-pay-url
```
##### Post Body & 
[Request Body](#faas-payment-through-bitgin-frontend)

<br />

- FaaS Get Receipt From BITGIN Backend
##### POST Method
```
/bitgin-receipt
```

##### Post Body
[Request Body](#faas-get-receipt-from-bitgin-backend)