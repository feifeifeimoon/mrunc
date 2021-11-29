package app

import (
	"github.com/feifeifeimoon/mrunc/pkg/container"
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "create <container-id> <path-to-bundle>",
		Long:    "create a container by id and bundle path",
		Example: "mrunc create busybox-1 bundle",
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id, bundle := args[0], args[1]
			_, _ = container.NewContainer(id, bundle)
		},
	}
}
