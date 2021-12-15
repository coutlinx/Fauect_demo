package Hander

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	mid "linx/Final_Project/Config"
)

func FreeLimit(c *gin.Context)  {
	linx :=mid.Linx{}
	client,err := mid.GetClient()
	if err != nil{
		respError(c,err)
		return
	}
	linx.AddressLinhao = c.PostForm("address")
	Address := common.HexToAddress(linx.AddressLinhao)
	if err != nil {
		respError(c, err)
		return
	}

	faucet, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		return
	}

	res, err := mid.FreedLimit(client,faucet,Address)
	if err != nil {
		respError(c,err)
		return
	}
	respOK(c,res)
	client.Close()
}