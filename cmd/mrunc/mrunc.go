package main

import (
	"fmt"
	"github.com/feifeifeimoon/mrunc/cmd/mrunc/app"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:  "mrunc ",
		Long: "mrunc is a container runtime implementation，comply with OCI specifications",
	}
	rootCmd.AddCommand(app.NewCreateCommand())
	rootCmd.AddCommand(app.NewVersionCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("execute err", err)
	}
}
