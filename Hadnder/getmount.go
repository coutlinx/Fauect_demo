package Hander

import (
	"github.com/gin-gonic/gin"
	mid "linx/Final_Project/Config"
)

func GetContractAmount(c *gin.Context)  {
	client := mid.GetClient()
	contract,err :=mid.GetCallFaceTx(*client)
	if err != nil{
		respError(c,err)
	}
	res:= mid.GetContractAmount(contract)
	respOK(c,res)
	client.Close()
}