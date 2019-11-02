package cmd

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

func ExecSelectCmd(c *cli.Context) error {
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

	if err := os.Chdir(path); err != nil {
		return err
	}

	files, err := readAllFiles(path)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	cmd := exec.Command(config.Selector)
	cmd.Stdin = strings.NewReader(strings.Join(files, "\n"))
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command(config.Editor, strings.TrimRight(buf.String(), "\n"))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func readAllFiles(path string) ([]string, error) {
	var paths []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return paths, err
	}
	return paths, nil
}
