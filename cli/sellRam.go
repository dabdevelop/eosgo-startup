package cli

import (
	"fmt"
	"encoding/json"
	
	"github.com/eoscanada/eos-go/system"
	"github.com/eoscanada/eos-go"
)

func SellRam(api *eos.API, accountName eos.AccountName, numBytes int32){
	action := system.NewSellRAM(accountName, uint64(numBytes))
	action.Account = eos.AN("eosio.token")
	actions := []*eos.Action{action}
	resp, err := api.SignPushActions(actions...)
	if err == nil{
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
	
}