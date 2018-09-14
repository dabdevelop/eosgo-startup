package cli

import (
	"encoding/json"
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/system"
)

func SellRam(api *eos.API, accountName eos.AccountName, numBytes int32) {
	action := system.NewSellRAM(accountName, uint64(numBytes))
	actions := []*eos.Action{action}
	resp, err := api.SignPushActions(actions...)
	if err == nil {
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}

}
