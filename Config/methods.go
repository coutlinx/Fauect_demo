package Config

import (
	Faucet "Final_Project/Faucet/GO"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

const (
	chainID       = 102400
	gasLimit      = uint64(21000)
	FaucetAddress = "0x25cC555dC2600b609DC6cA27C3b4E25C38d4C13A"
)

// GetClient 获取 ethclient.Client 对象
func GetClient() (*ethclient.Client, error) {
	rpcDial, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		return nil, err
	}
	client := ethclient.NewClient(rpcDial)
	return client, nil
}

// GetFacetTx 获取 Faucet.ApinameTransactor 对象
func GetFacetTx(client *ethclient.Client) (*Faucet.ApinameTransactor, error) {
	contract, err := Faucet.NewApinameTransactor(common.HexToAddress(FaucetAddress), client)
	return contract, err
}

// GetCallFaceTx 获取 Faucet.ApinameCaller 对象
func GetCallFaceTx(client ethclient.Client) (*Faucet.ApinameCaller, error) {
	contract, err := Faucet.NewApinameCaller(common.HexToAddress(FaucetAddress), &client)
	return contract, err
}

// 初步获取 *bind.TransactOpts 对象
func setopts(client *ethclient.Client, privateKey *ecdsa.PrivateKey, address common.Address, value int64) *bind.TransactOpts {
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
	auth.Value = big.NewInt(value)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth
}

// GetTxopts 获取最终 *bind.TransactOpts 对象
func GetTxopts(client *ethclient.Client, value int64) *bind.TransactOpts {
	privateKey, publicKey := GetPublicKeyPrivateKey()
	opts := setopts(client, privateKey, publicKey, value)
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
	PK := "b4b9cefb7c5a83e98d40a7933ee79730b5486e33d2402fa0889fedb65b1a2e60"
	privateKey, err := crypto.HexToECDSA(PK)
	if err != nil {
		panic(err)
	}
	publicKey := common.HexToAddress("0x155b091aEe4b10D2f995f21ec6bE3435A2c4e1E9")
	return privateKey, publicKey
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
func SendETH(client *ethclient.Client, contract *Faucet.ApinameTransactor, Value *big.Int, Address common.Address) (*types.Transaction, error) {
	opts := GetTxopts(client, 0)
	res, err := contract.Withdraw(opts, Value, Address)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetContractAmount 调用 Faucet合约 查看合约ETH数目
func GetContractAmount(contract *Faucet.ApinameCaller) (*big.Int, error) {

	res, err := contract.GetContractAmount(nil)
	if err != nil {
		return big.NewInt(0), err
	}
	return res, nil
}

// GetLimit 调用合约得到用户的提取ETH数
func GetLimit(contract *Faucet.ApinameCaller, address common.Address) (*big.Int, error) {
	res, err := contract.Limit(nil, address)
	if err != nil {
		return big.NewInt(0), err
	}
	return res, nil
}

func SendALl(client *ethclient.Client, contract *Faucet.ApinameTransactor, NewAddress common.Address) (*types.Transaction, error) {
	PK := "c28ae2f0b29deeba2aec5b6d4fe91965f5f3bd74e80ff406eb6c75f526f902a7"
	privateKey, err := crypto.HexToECDSA(PK)
	opts := setopts(client, privateKey, common.HexToAddress("0xA7c634F44076e13d025F24D6f4C8E1F51fAd400b"), 0)
	res, err := contract.TransferOther(opts, NewAddress)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Charge(client *ethclient.Client, contract *Faucet.ApinameTransactor, value int64) (*types.Transaction, error) {
	opts := GetTxopts(client, value)
	res, err := contract.Charge(opts)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FreedLimit(client *ethclient.Client, contract *Faucet.ApinameTransactor, Address common.Address) (*types.Transaction, error) {
	opts := GetTxopts(client, 0)
	res, err := contract.EmptyLimit(opts, Address)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FreeLimits() {
	fmt.Println("start")
	client, err := GetClient()
	if err != nil {
		panic(err)
		return
	}
	contract, err := GetFacetTx(client)
	if err != nil {
		panic(err)
		return
	}
	address, err := SelectAllAddress()
	if err != nil {
		panic(err)
	}
	for _, add := range address {
		_, err := FreedLimit(client, contract, add)
		if err != nil {
			panic(err)
			return
		}
	}
}
