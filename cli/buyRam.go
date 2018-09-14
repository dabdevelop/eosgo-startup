package cli

import (
	"fmt"
	"encoding/json"
	
	"github.com/eoscanada/eos-go/system"
	"github.com/eoscanada/eos-go"
)

func BuyRam(api *eos.API, payer, receiver eos.AccountName, numBytes int){
	action := system.NewBuyRAMBytes(payer, receiver, uint32(numBytes))
	actions := []*eos.Action{action}
	resp, err := api.SignPushActions(actions...)
	if err == nil{
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
	
}