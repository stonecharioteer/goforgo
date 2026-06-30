package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	selfUpdateYes   bool
	selfUpdateCheck bool
)

var selfUpdateCmd = &cobra.Command{
	Use:   "self-update",
	Short: "Check for and install the latest GoForGo version",
	Long: `Checks the latest GitHub tag and optionally updates the installed goforgo binary.

By default, this command asks for confirmation before updating.
Use --yes to skip the prompt.
Use --check to only check and print status.`,
	RunE: runSelfUpdate,
}

func runSelfUpdate(cmd *cobra.Command, args []string) error {
	latest, isNewer, err := checkForUpdate(version, defaultTagsURLForTests, nil)
	if err != nil {
		return fmt.Errorf("failed to check for updates: %w", err)
	}

	if !isNewer {
		return writeCommandOutput(cmd, "✅ goforgo is up to date (%s)\n", version)
	}

	if err := writeCommandOutput(cmd, "🔔 Update available: %s (current: %s)\n", latest, version); err != nil {
		return err
	}
	if err := writeCommandOutput(cmd, "   Will run: %s\n", updateInstallCmd); err != nil {
		return err
	}

	if selfUpdateCheck {
		return nil
	}

	if !selfUpdateYes {
		ok, err := askForConfirmation(cmd)
		if err != nil {
			return err
		}
		if !ok {
			return writeCommandOutput(cmd, "Update cancelled.\n")
		}
	}

	if err := writeCommandOutput(cmd, "Updating goforgo...\n"); err != nil {
		return err
	}
	goCmd := exec.Command("go", "install", "github.com/stonecharioteer/goforgo/cmd/goforgo@latest")
	goCmd.Stdout = cmd.OutOrStdout()
	goCmd.Stderr = cmd.ErrOrStderr()

	if err := goCmd.Run(); err != nil {
		return fmt.Errorf("self-update failed: %w", err)
	}

	return writeCommandOutput(cmd, "✅ Update complete. Run 'goforgo --version' to verify.\n")
}

func askForConfirmation(cmd *cobra.Command) (bool, error) {
	if err := writeCommandOutput(cmd, "Proceed with update? [y/N]: "); err != nil {
		return false, err
	}

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("failed to read confirmation: %w", err)
	}

	answer := strings.TrimSpace(strings.ToLower(line))
	return answer == "y" || answer == "yes", nil
}

func writeCommandOutput(cmd *cobra.Command, format string, args ...any) error {
	if _, err := fmt.Fprintf(cmd.OutOrStdout(), format, args...); err != nil {
		return fmt.Errorf("failed to write command output: %w", err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(selfUpdateCmd)
	selfUpdateCmd.Flags().BoolVarP(&selfUpdateYes, "yes", "y", false, "run update without confirmation prompt")
	selfUpdateCmd.Flags().BoolVar(&selfUpdateCheck, "check", false, "only check for updates, do not install")
}
