package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"math/rand"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate random passwords",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Length of Password")
	generateCmd.Flags().BoolP("digits", "d", false, "Include Digits")
	generateCmd.Flags().BoolP("special-chars", "s", false, "Include Special Characters")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if isDigits {
		charSet += "0123456789"
	}

	if isSpecialChars {
		charSet += "!@#$%^&*()_+[]{}|;:,.<>?-="
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charSet[rand.Intn(len(charSet))]
	}

	fmt.Println("Generating Password...")
	fmt.Println(string(password))
}