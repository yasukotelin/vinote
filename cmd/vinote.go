package cmd

import (
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/urfave/cli"
)

// ExecVinoteCmd runs on the user specified note directory
// with user specified editor (vi, Vim or Neovim).
func ExecVinoteCmd(c *cli.Context) error {
	path, err := replaceHomeDirectory("~/Dropbox/note")
	if err != nil {
		return err
	}


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

func replaceHomeDirectory(path string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	newPath := strings.Replace(path, "~", usr.HomeDir, 1)
	return newPath, nil
}
