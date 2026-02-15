// XebecCorporation.Dots - CLI for XEBEC CORPORATION ecosystem
package main

import (
	"fmt"
	"os"

	"github.com/XebecCorporation/XebecCorporation.Dots/cmd/xebec/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
