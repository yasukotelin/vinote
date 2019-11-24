package cmd

import (
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

// ExecVinoteCmd runs on the user specified note directory
// with user specified editor (vi, Vim or Neovim).
func ExecVinoteCmd(c *cli.Context) error {
	config, err := ReadConfig()
	if err != nil {
		return err
	}

	path, err := config.GetFullPath()
	if err != nil {
		return err
	}

	// Change the current directory to note directory.
	if err := os.Chdir(path); err != nil {
		return err
	}

	cmd := exec.Command(config.Editor, ".")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdin
	cmd.Stdin = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
