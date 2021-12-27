package Hander

import (
	mid "Final_Project/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Chager(c *gin.Context) {
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
	if err != nil {
		respError(c, err)
		return
	}
	linx.AmountLinhao, err = strconv.ParseInt(c.PostForm("amount"), 10, 64)
	if err != nil {
		respError(c, err)
		return
	}
	fmt.Println(linx.AmountLinhao)

	res, err := mid.Charge(client, contract, linx.AmountLinhao)
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, res)
	client.Close()
}
