package cmd

import (
	"github.com/MingmarGyalzenSherpa44/chat-app/internal/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a server",
	Run: func(cmd *cobra.Command, args []string) {

		server.InitServer()

	},
}
