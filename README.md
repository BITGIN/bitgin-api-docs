# Table of contents
- [API Authentication](#api-authentication)

- [Fiat-as-a-Service](#fiat-as-a-service)
  - [Services](#services)
    - [FaaS Payment Through BITGIN Frontend](#faas-payment-through-bitgin-frontend)
    - [FaaS Get Receipt From BITGIN Backend](#faas-get-receipt-from-bitgin-backend)

- [Mine Share Service](#mine-share-service)
  - [REST API](#rest-api)
    - [Query BITGIN Addresses](#query-bitgin-addresses)
    - [Mine Share](#mine-share)

- [Mock Server](#mock-server)
  - [How to compile](#how-to-compile)
  - [Setup configuration](#setup-configuration)
  - [REST API of Mock Server](#rest-api-of-mock-server)
	

# API Authentication
- The first things you need to use the FaaS is your API key and your secret key.
- You can use API key and secret key to generate the necessary parameters that contain  `SIGN` , `KEY` , `NONCE`

### How to sign your data ?
- Please refer to the code snippet on the github project to know how to sign your data.
	- [Go](https://github.com/BITGIN/bitgin-api-docs/blob/main/handler/handler.go#L24)

##### [Back to top](#table-of-contents)
# Fiat-as-a-Service

- The quickest way to build conversion-optimized cryptocurrency payment forms, hosted on BITGIN.

- Fiat-as-a-Service creates a secure, BITGIN-hosted payment page that lets you collect payments quickly. It works across cryptocurrency and can help increase your conversion. Fiat-as-a-Service makes it easy to build a first-class payments experience.

### How to contact BITGIN ? 
- [Contact us](https://bitgin.freshdesk.com/support/tickets/new)
### Try it now
- Use [Mock Server](#mock-server) to test BITGIN Fiat-as-a-Service right away.
  - Step 1: Deploy the Mock Server
  - Step 2: Call [FaaS Payment Through BITGIN Frontend](#faas-payment-through-bitgin-frontend) to get BITGIN Frontend URL
  - Step 3: Open the URL then complete payment
  - Step 4: Call [FaaS Get Receipt From BITGIN Backend](#faas-get-receipt-from-bitgin-backend) to acquire payment receipts

##### [Back to top](#table-of-contents)



# Services

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

##### Headers
| Key | Value | Note |
| :---  | :--- | :---        |
| Content-Type | application/json | required, JSON Type |


##### Body Format

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
  "url": "https://bitgin.net/fiat-as-a-service?sign=<sign>&key=<key>&nonce=<nonce>&body=<body>"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| url | string |  URL is the payment site provided by BITGIN|

##### [Back to top](#table-of-contents)


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

##### Headers
| Key | Value | Note |
| :---  | :--- | :---        |
| Content-Type | application/json | required, JSON Type |
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
| currency | string | Currency of payment (USDT, ETH, Bitcoin...)|
| fee_currency | string | Currency of fee of payment (USDT, ETH, Bitcoin...)|
| address | string |  Deposit address of payment|
| chain | string |  Chain of payment (Tron, Ethereum, Bitcoin...)|
| tx_id | string | Transaction ID is a string that identifies a specific transaction on the blockchain|
| is_deduction | boolean | |


"authorization_error"
"invalid_request_format"
"invalid_withdrawal_value"
"unknown_error"
"unexpected_error"
"invalid_chain_type"
"invalid_currency_type"
"invalid_api_request"
"permission_denied"

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

##### [Back to top](#table-of-contents)


# REST API
## Query BITGIN Addresses


##### Request

**POST** `BITGIN_DOMAIN`/mine/v1/query

##### Request Format

An example of the request:

###### API

```
/mine/v1/query
```
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
| currency | string | required (ETH, USDT, Bitcoin...) | Specify currency type |
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

**POST** `BITGIN_DOMAIN`/mine/v1/share



##### Request Format

An example of the request:

###### API

```
/mine/v1/share
```

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
### FaaS Payment Through BITGIN Frontend
```
http://localhost:8888/faas/v1/pay
```

[API definition](#faas-payment-through-bitgin-frontend)


### FaaS Get Receipt From BITGIN Backend

```
http://localhost:8888/faas/v1/receipt
```

[API definition](#faas-get-receipt-from-bitgin-backend)

<br />

- ## Mine Share Service

### Query BITGIN Addresses

```
http://localhost:8888/mine/v1/query
```

[API definition](#query-bitgin-addresses)

### Mine Share
```
http://localhost:8888/mine/v1/share
```
[API definition](#mine-share)

##### [Back to top](#table-of-contents)
