package Config

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	Faucet "linx/Final_Project/Faucet/GO"
	Hander "linx/Final_Project/Hadnder"
	"math/big"
)
const(
	chainID       = 102400
	gasLimit      = uint64(21000)
	FaucetAddress = "0x3e33B941387752a16E62C9e5C059F03f4dD71F3b"
)

// GetClient 获取 ethclient.Client 对象
func GetClient()  *ethclient.Client{
	rpcDial, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	client := ethclient.NewClient(rpcDial)
	return client
}

// GetFacetTx 获取 Faucet.ApinameTransactor 对象
func GetFacetTx(client *ethclient.Client)  (* Faucet.ApinameTransactor, error){
	contract, err := Faucet.NewApinameTransactor(common.HexToAddress(FaucetAddress), client)
	return contract,err
}

// GetCallFaceTx 获取 Faucet.ApinameCaller 对象
func GetCallFaceTx(client ethclient.Client)   (* Faucet.ApinameCaller, error){
	contract,err := Faucet.NewApinameCaller(common.HexToAddress(FaucetAddress), &client)
	return contract,err
}

// 初步获取 *bind.TransactOpts 对象
func setopts(client *ethclient.Client, privateKey *ecdsa.PrivateKey, address common.Address) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		panic(err)
	}
	nonce, err := Getnonce(client, address)
	if err != nil {
		panic(err)
	}
	gasPrice, err := GetgasPrice(client)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth
}

// SetTxData 获取 TXData 结构体
func SetTxData(client *ethclient.Client, nonce uint64, address *common.Address, value *big.Int) (*types.Transaction, error) {
	gasPrice, err := GetgasPrice(client)
	if err != nil {
		panic(err)
	}
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       address,
		Value:    value,
	})
	return tx, nil
}

// GetTxopts 获取最终 *bind.TransactOpts 对象
func GetTxopts(client *ethclient.Client) *bind.TransactOpts{
	privateKey, publicKey := GetPublicKeyPrivateKey()
	opts := setopts(client, privateKey, publicKey)
	return opts
}

// Getnonce 获取 nonce
func Getnonce(client *ethclient.Client, address common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, err
	} else {
		return nonce, nil
	}

}

// GetPublicKeyPrivateKey 获取公钥和私钥
func GetPublicKeyPrivateKey() (*ecdsa.PrivateKey, common.Address) {
	const file = "Final_Project/Chain/Linx/data/keystore/UTC--2021-12-14T00-41-15.258781500Z--513aba4e1b4eddb77196c31df9ed9fedf914b86c"
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	passwd := "102400"
	key, err := keystore.DecryptKey(jsonBytes, passwd)
	if err != nil {
		panic(err)
	}
	return key.PrivateKey, key.Address
}

// GetgasPrice 获取 gasPrice
func GetgasPrice(client *ethclient.Client) (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return big.NewInt(0), err
	} else {
		return gasPrice, nil
	}

}

// SendETH 调用 Faucet合约 发送ETH
func SendETH(client *ethclient.Client,contract *Faucet.ApinameTransactor,tx *Hander.Tx) *types.Transaction{
	opts:=GetTxopts(client)
	res, err := contract.Withdraw(opts, tx.Value , tx.Address)
	if err != nil {
		panic(err)
	}
	return res
}

// GetContractAmount 调用 Faucet合约 查看合约ETH数目
func GetContractAmount(contract *Faucet.ApinameCaller)  *big.Int{
	res,err:=contract.GetContractAmount(nil)
	if err != nil {
		panic(err)
	}
	return res
}

// GetLimit 调用合约得到用户的提取ETH数
func GetLimit(contract *Faucet.ApinameCaller,address common.Address)  *big.Int{
	res,err := contract.Limit(nil,address)
	if err != nil{
		panic(err)
	}
	return res
}