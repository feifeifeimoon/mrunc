package app

import (
	"fmt"
	"github.com/feifeifeimoon/mrunc/pkg/version"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Example: "mrunc version",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.Get())
		},
	}
}
