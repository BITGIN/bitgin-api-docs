# Table of contents

- [Get Started](#get-started)
- [API Authentication](#api-authentication)
- [Services](#services)
    - [FaaS Payment Through Bitgin Frontend](#faas-payment-through-bitgin-frontend)
    - [FaaS Get Receipt From Bitgin Backend](#faas-get-receipt-from-bitgin-backend)
- [Mock Server](#mock-server)

	


# Get Started




##### [Back to top](#table-of-contents)


# API Authentication
- The first things you need to use the FaaS is your API key and your secret key.
- You can use API key and secret key to generate the necessary parameters that contain  `<sign>` , `<key>` , `<nonce>`
- Please refer to the code snippet on the github project to know how to sign your data.
	- [Go](https://github.com/BITGIN/faas-api-docs/blob/main/main.go#L103)

##### [Back to top](#table-of-contents)

<br />
<br />

# Services

<br />


## FaaS Payment Through Bitgin Frontend

##### URL
```
https://bitgin.net/fiat-as-a-service?sign=<sign>&key=<key>&nonce=<nonce>&body=<body>
```

The URL includes the following parameters:

##### URL Parameters

| Field | Note | Description |
| :---  | :--- | :---        |
| sign | acquired from Bitgin | Which is the SHA256 HMAC of the following four strings, using your API secret, as a hex string: Request timestamp (same as above), HTTP method in uppercase (e.g. GET or POST), Request path, including leading slash and any URL parameters but not including the hostname|
| key  | acquired from Bitgin | The key string for Bitgin verification merchant |
| nonce | represented by milliseconds |  Unix time of current time, the number of `milliseconds` elapsed since January 1, 1970 UTC |
| body | base64 encoding as specified by RFC 4648 | Specify the base64 encoded payment information [Body Format](#body-format)|



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

##### [Back to top](#table-of-contents)


<br />
<br />

## FaaS Get Receipt From Bitgin Backend

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
| data | JSON | Receipts data |
| request_id | string | Request id For Debug  |



##### [Back to top](#table-of-contents)

<br />
<br />

# Mock Server

### How to compile

- Clone project
```
$ git clone https://github.com/BITGIN/faas-api-docs.git
```
##### How to get start with [Go](https://go.dev/doc/tutorial/getting-started)


- run mock server on `port :8881`
```
$ go run main.go
```


### Setup configuration

>	NOTE: Configure in main.go
```go
key    = "<API_KEY>"  // API_KEY acquired from BITGIN
secret = "<SECRET_KEY>" // SECRET_KEY acquired from BITGIN
```
<br />

### Services

- FaaS Payment Through Bitgin Frontend
##### POST Method
```
http://localhost:8881/bitgin-pay-url
```

<br />

- FaaS Get Receipt From Bitgin Backend
##### POST Method
```
http://localhost:8881/bitgin-receipt
```