package cli

import (
	"fmt"
	"encoding/json"

	"github.com/eoscanada/eos-go/token"
	"github.com/eoscanada/eos-go"
)

func Transfer(api *eos.API, from eos.AccountName, to eos.AccountName, quantity eos.Asset, memo string, code eos.AccountName){
	action := token.NewTransfer(from, to, quantity, memo)
	action.Account = code
	actions := []*eos.Action{action}
	resp, err := api.SignPushActions(actions...)
	if err == nil{
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
	
}