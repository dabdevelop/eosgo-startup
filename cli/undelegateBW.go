package cli

import (
	"fmt"
	"encoding/json"

	"github.com/eoscanada/eos-go/system"
	"github.com/eoscanada/eos-go"
)

func UndelegateBW(api *eos.API, from, receiver eos.AccountName, cpuStake, netStake eos.Asset){
	actions := []*eos.Action{}
	actions = append(actions, system.NewUndelegateBW(from, receiver, cpuStake, netStake))
	resp, err := api.SignPushActions(actions...)
	if err == nil{
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}