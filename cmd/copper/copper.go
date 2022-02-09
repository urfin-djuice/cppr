package main

import (
	"copper/internal/pkg/copperclient"
	"fmt"
)

func main() {
	c := copperclient.NewCopperClient()

	dt, err := c.CreateDepositTarget(copperclient.CurrencySOL)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	} else {
		fmt.Printf("%+v\n", *dt)
	}

}
