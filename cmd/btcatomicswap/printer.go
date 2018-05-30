package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string
}

//PrintProperty is a property to print in json or human readable
type PrintProperty struct {
	Name         string
	FriendlyName string
	Value        string
}

var (
	ppSecret              = PrintProperty{Name: "Secret", FriendlyName: "Secret"}
	ppHash                = PrintProperty{Name: "Hash", FriendlyName: "Hash"}
	ppContractFee         = PrintProperty{Name: "ContractFee", FriendlyName: "Contract fee"}
	ppRefundFee           = PrintProperty{Name: "RefundFee", FriendlyName: "Refund fee"}
	ppRedeemFee           = PrintProperty{Name: "RedeemFee", FriendlyName: "Redeem fee"}
	ppRedeemTransaction   = PrintProperty{Name: "RedeemTransaction", FriendlyName: "Redeem transaction"}
	ppContract            = PrintProperty{Name: "Contract", FriendlyName: "Contract"}
	ppContractValue       = PrintProperty{Name: "ContractValue", FriendlyName: "ContractValue"}
	ppSecretHash          = PrintProperty{Name: "SecretHash", FriendlyName: "Secret hash"}
	ppContractTransaction = PrintProperty{Name: "ContractTransaction", FriendlyName: "Contract transaction"}
	ppRefundTransaction   = PrintProperty{Name: "RefundTransaction", FriendlyName: "Refund transaction"}
	ppRecipientAddress    = PrintProperty{Name: "RecipientAddress", FriendlyName: "Recipient address"}
	ppLockTime            = PrintProperty{Name: "LockTime", FriendlyName: "Lock time"}
	ppRefundAddress       = PrintProperty{Name: "RefundAddress", FriendlyName: "Refund address"}
)

func getPrintProperty(property PrintProperty, value interface{}) PrintProperty {
	stringVal := fmt.Sprintf("%v", value)
	property.Value = stringVal
	return property
}

func print(actionName string, actionDescription string, printProperties []PrintProperty) {
	if *automatedFlag {

		m := map[string]interface{}{
			"action": actionName,
			"data":   printProperties,
		}

		jsonPrint(m)

	} else { //pretty print
		fmt.Println(actionDescription)

		for _, printProperty := range printProperties {
			fmt.Printf("%s: %s", printProperty.FriendlyName, printProperty.Value)
		}

	}
}
func jsonPrint(m interface{}) {
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

//PrintInfo prints extra info, only for non automatic swaps
func PrintInfo(args ...interface{}) {
	if *automatedFlag {
		fmt.Fprintln(os.Stdout, args...)
	}

}
