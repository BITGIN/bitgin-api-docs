package main

import (
	"github.com/BITGIN/bitgin-api-docs/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/v1/faas/pay", handler.FaasPayHandler)

	e.GET("/v1/faas/receipt", handler.FaasReceiptHandler)

	e.POST("/v1/mine/query", handler.MineQueryAddressesHandler)

	e.POST("/v1/mine/share", handler.MineShareHandler)

	e.Logger.Fatal(e.Start(":8888"))
}
