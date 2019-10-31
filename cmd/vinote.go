package cmd

import (
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func ExecVinoteCmd(c *cli.Context) error {
	path := "~/Dropbox/note"

	// Change the current directory to note directory.
	if err := os.Chdir(path); err != nil {
		return err
	}

	cmd := exec.Command("nvim")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdin
	cmd.Stdin = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
