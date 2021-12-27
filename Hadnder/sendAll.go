package Hander

import (
	mid "Final_Project/Config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func SendAll(c *gin.Context) {
	linx := mid.Linx{}
	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		return
	}

	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		return
	}
	linx.AddressLinhao = c.PostForm("address")
	res, err := mid.SendALl(client, contract, common.HexToAddress(linx.AddressLinhao))
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, res)
	client.Close()
}
