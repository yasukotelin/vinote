package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

type Config struct {
	Editor   string `json:"editor"`
	Path     string `json:"path"`
	Selector string `json:"selector"`
}

func (c *Config) GetFullPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	newPath := strings.Replace(c.Path, "~", usr.HomeDir, 1)
	return newPath, nil
}

// ExecConfigCmd opens config json file `.vinote.json` on the home direcotory.
// If it still donesn't exist, this will create it.
func ExecConfigCmd(c *cli.Context) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	editor := ""

	config, err := readConfig(path)
	if err == nil {
		editor = config.Editor
	} else {
		editor = askWhatUseEditor()
	}

	cmd := exec.Command(editor, path)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdin
	cmd.Stdin = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func getConfigPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, ".vinote.json"), nil
}

func readConfig(path string) (*Config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func ReadConfig() (*Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	config, err := readConfig(path)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func askWhatUseEditor() string {
	var editor string
	fmt.Println("What do you use the editor?")
	fmt.Print("(vi or vim or nvim) > ")
	fmt.Scanln(&editor)
	return editor
}
