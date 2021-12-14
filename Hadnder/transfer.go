package Hander

import (
	mid "linx/Final_Project/Config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"strconv"
)

// Tx 定义 交易结构体
type Tx struct {
	Address common.Address
	Value *big.Int
}

// Transfer 路由中间件 Transfer
func Transfer(c *gin.Context) {
	client := mid.GetClient()
	Address := common.HexToAddress(c.PostForm("ADDRESS"))
	value,err:=strconv.ParseInt(c.PostForm("VALUE"), 10, 64)
	if err!=nil{
		respError(c, err)
		return
	}
	Value := big.NewInt(value)
	tx:= Tx{
		Address:Address,
		Value: Value,
	}
	faucet,err:=mid.GetFacetTx(client)
	if err!=nil{
		panic(err)
	}
	 res :=mid.SendETH(client,faucet,&tx)
	respOK(c,res)
	 client.Close()
}
