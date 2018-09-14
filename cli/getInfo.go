package cli

import (
	"fmt"
	"encoding/json"

	"github.com/eoscanada/eos-go"
)

func GetInfo(api *eos.API){
	info, err := api.GetInfo()
	if err == nil {
		data, _ := json.MarshalIndent(info, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println("Request failed")
	}
}