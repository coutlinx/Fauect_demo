package Hander

import (
	mid "Final_Project/Config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func Limit(c *gin.Context) {
	linx := mid.Linx{}
	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		return
	}
	linx.AddressLinhao = c.PostForm("address")
	address := common.HexToAddress(linx.AddressLinhao)
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		return
	}
	res, err := mid.GetLimit(contract, address)
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, res)
	client.Close()
}
