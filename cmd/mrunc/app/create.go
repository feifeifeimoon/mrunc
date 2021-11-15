package app

import (
	"github.com/feifeifeimoon/mrunc/pkg/container"
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "create <container-id> <path-to-bundle>",
		Long:    "create a container by id and bundle path",
		Example: "create ae86 ",
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, bundle := args[0], args[1]
			_, err := container.NewContainer(id, bundle)
			return err
		},
	}
}
