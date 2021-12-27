package Hander

import (
	mid "Final_Project/Config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"strconv"
	"time"
)

var AT = make(map[string]int)

// Transfer 路由中间件 Transfer
func Transfer(c *gin.Context) {
	linx := mid.Linx{}
	linx.AddressLinhao = c.PostForm("address")
	if AT[linx.AddressLinhao]-time.Now().Second() < 5 && AT[linx.AddressLinhao]-time.Now().Second() > -5 && AT[linx.AddressLinhao] != 0 {
		respError(c, "请过几秒在试下吧")
		return
	}
	AT[linx.AddressLinhao] = time.Now().Second()
	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		return
	}
	linx.AmountLinhao, err = strconv.ParseInt(c.PostForm("amount"), 10, 64)
	if err != nil {
		respError(c, err)

		return
	}
	Address := common.HexToAddress(linx.AddressLinhao)
	if err != nil {
		respError(c, err)

		return
	}

	Value := big.NewInt(linx.AmountLinhao)

	faucet, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		return
	}
	res, err := mid.SendETH(client, faucet, Value, Address)
	if err != nil {
		respError(c, err)
		return
	}
	err = mid.InsertAdd(linx.AddressLinhao, c.ClientIP(), linx.AmountLinhao)
	if err != nil {
		respError(c, err)
		return
	}
	err = mid.InsertCon(linx.AmountLinhao)
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, res)
	client.Close()
}
