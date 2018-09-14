package cli

import (
	"fmt"
	"encoding/json"
	
	"github.com/eoscanada/eos-go"
)

func GetAccount(api *eos.API, account eos.AccountName){
	resp, _ := api.GetAccount(account)
	data, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(data))
}