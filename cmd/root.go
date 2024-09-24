package cmd

import (
	"fmt"
	"os"

	"github.com/MingmarGyalzenSherpa44/chat-app/internal/database"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-chat-app",
	Short: "A CLI chat application",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	database.Connect()
}
