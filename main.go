package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{
		Use:  "mrunc ",
		Long: "mrunc is a container runtime implementation，comply with OCI specifications",
	}
	rootCmd.AddCommand(NewCreateCommand())
	rootCmd.AddCommand(NewVersionCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("execute err", err)
	}
}
