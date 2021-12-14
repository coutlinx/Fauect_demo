package Hander

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	mid "linx/Final_Project/Config"
)

func Limit(c *gin.Context)  {
	client := mid.GetClient()
	address := common.HexToAddress(c.PostForm("ADDRESS"))
	contract,err := mid.GetCallFaceTx(*client)
	if err != nil{
		respError(c,err)
	}
	res :=mid.GetLimit(contract,address)
	respOK(c,res)
}
