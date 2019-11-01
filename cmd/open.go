package cmd

import (
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func ExecOpenCmd(c *cli.Context) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	config, err := ReadConfig(path)
	if err != nil {
		return err
	}

	path, err = config.getFullPath()
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
