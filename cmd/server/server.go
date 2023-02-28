package server

import (
	"github.com/loft-sh/devpod/cmd/flags"
	"github.com/spf13/cobra"
)

// NewServerCmd returns a new root command
func NewServerCmd(flags *flags.GlobalFlags) *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "DevPod Server commands",
	}

	serverCmd.AddCommand(NewListCmd(flags))
	serverCmd.AddCommand(NewSSHCmd(flags))
	serverCmd.AddCommand(NewStopCmd(flags))
	serverCmd.AddCommand(NewStartCmd(flags))
	serverCmd.AddCommand(NewStatusCmd(flags))
	return serverCmd
}