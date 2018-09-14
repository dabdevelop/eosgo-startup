package main

import (
	"./cli"

	"github.com/eoscanada/eos-go"
)

var (
	api *eos.API
)

func init(){
	rpc := "https://api.eosnewyork.io"
	key := "5KdB3AtHrJFZez87KHAQej4ucZ7NAwN2VenB1D9fVDSMdGFq4jK"
	api = eos.New(rpc)
	signer := eos.NewKeyBag()
	signer.ImportPrivateKey(key)
	api.SetSigner(signer)
}

func main(){
	transfer()
}

func transfer(){
	from := eos.AN("xxxxxaccount")
	to := eos.AN("orangeisluck")
	quantity, _ := eos.NewEOSAssetFromString("0.0001 EOS")
	memo := ""
	code := eos.AN("eosio.token")
	cli.Transfer(api, from, to, quantity, memo, code)
}

func buyRam(){
	cli.BuyRam(api, eos.AN("xxxxxaccount"), eos.AN("xxxxxaccount"), 1 * 1024)
}

func sellRam(){
	cli.SellRam(api, eos.AN("xxxxxaccount"), 1 * 1024)
}

func delegateBW(){
	from := eos.AN("xxxxxaccount")
	receiver := eos.AN("xxxxxaccount")
	cpuStake, _ := eos.NewAsset("0.0001 EOS")
	netStake, _ := eos.NewAsset("0.0001 EOS")
	cli.DelegateBW(api, from, receiver, cpuStake, netStake, false)
}

func undelegateBW(){
	from := eos.AN("xxxxxaccount")
	receiver := eos.AN("xxxxxaccount")
	cpuStake, _ := eos.NewAsset("0.0001 EOS")
	netStake, _ := eos.NewAsset("0.0001 EOS")
	cli.UndelegateBW(api, from, receiver, cpuStake, netStake)
}


func newKeyPair(){
	cli.NewKeyPair()
}

func getCurrencyBalance(){
	account := eos.AN("xxxxxaccount")
	symbol := "EOS"
	tokenAccount := eos.AN("eosio.token") 
	cli.GetCurrencyBalance(api,account, symbol, tokenAccount)
}

func getInfo(){
	cli.GetInfo(api)
}

func getAccount(){
	cli.GetAccount(api, eos.AN("xxxxxaccount"))
}