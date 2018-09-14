package cli

import (
	"encoding/json"
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/system"
)

func DelegateBW(api *eos.API, from, receiver eos.AccountName, cpuStake, netStake eos.Asset, doTransfer bool) {
	actions := []*eos.Action{}
	actions = append(actions, system.NewDelegateBW(from, receiver, cpuStake, netStake, doTransfer))
	resp, err := api.SignPushActions(actions...)
	if err == nil {
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}
