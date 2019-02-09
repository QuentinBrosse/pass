package cmd

import (
	"fmt"
	"github.com/QuentinBrosse/pass/internal/app/pass/vault"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "Create a new vault",
	Run:   runCreate,
}

// Construct vault file path from the CMD-specified configuration directory and filename
func getVaultFilePath() string {
	return filepath.Join(PersistentFlags.ConfDirPath, PersistentFlags.VaultFilename)
}

// Create the specified vault file if it doesn't exist, exit otherwise
func runCreate(cmd *cobra.Command, args []string) {
	vaultFilePath := getVaultFilePath()

	if _, err := os.Stat(vaultFilePath); os.IsNotExist(err) {
		v := vault.NewVault(vaultFilePath, vault.DevEncryptionKey)
		v.Save()
	} else {
		fmt.Printf("Vault %s already exists", vaultFilePath)
		os.Exit(1)
	}
}
