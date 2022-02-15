package main

import (
	"github.com/BITGIN/bitgin-api-docs/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/faas/v1/pay", handler.FaasPayHandler)

	e.POST("/faas/v1/receipt", handler.FaasReceiptHandler)

	e.POST("/mine/v1/query", handler.MineQueryAddressesHandler)

	e.POST("/mine/v1/share", handler.MineShareHandler)

	e.Logger.Fatal(e.Start(":8888"))
}
