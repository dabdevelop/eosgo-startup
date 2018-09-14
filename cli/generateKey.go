package cli

import(
	"fmt"
	
	"github.com/eoscanada/eos-go/ecc"
)

func NewKeyPair(){
	priKey, _ := ecc.NewRandomPrivateKey()
	pubKey := priKey.PublicKey()
	fmt.Printf("Private Key: %s\nPublic  Key: %s\n", priKey, pubKey)
}