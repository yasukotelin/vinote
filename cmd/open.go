package cmd

import (
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func ExecOpenCmd(c *cli.Context) error {
	config, err := ReadConfig()
	if err != nil {
		return err
	}

	path, err := config.GetFullPath()
	if err != nil {
		return err
	}

	cmd := exec.Command("open", path)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
