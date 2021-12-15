package Hander

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mid "linx/Final_Project/Config"
)

func GetContractAmount(c *gin.Context)  {
	client,err := mid.GetClient()
	if err != nil{
		respError(c,err)
		fmt.Println(err)
		return
	}
	contract,err :=mid.GetCallFaceTx(*client)
	if err != nil{
		respError(c,err)
		fmt.Println(err)
		return
	}
	res,err:= mid.GetContractAmount(contract)
	if err != nil{
		respError(c,err)
		fmt.Println(err)
		return
	}
	respOK(c,res)
	client.Close()
}