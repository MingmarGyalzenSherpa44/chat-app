package cmd

import (
	"fmt"
	"log"

	"github.com/MingmarGyalzenSherpa44/chat-app/internal/database"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {
		email := args[0]
		username := args[1]
		password := args[2]

		fmt.Printf("Registering a new user!\n")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
			log.Fatal("Error while hashing password")
		}

		err = database.RegisterUser(email, username, string(hashedPassword))
		if err != nil {
			log.Fatal("Error registering a new user")
		}

		fmt.Printf("User registered successfully!")
	},
}
