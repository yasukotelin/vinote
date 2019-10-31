package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli"
)

type Config struct {
	Editor   string `json:"editor"`
	Path     string `json:"path"`
	Selector string `json:"selector"`
}

// ExecConfigCmd opens config json file `.vinote.json` on the home direcotory.
// If it still donesn't exist, this will create it.
func ExecConfigCmd(c *cli.Context) error {
	path, err := getVinoteJsonPath()
	if err != nil {
		return err
	}

	editor := ""

	if existsConfigFile(path) {
		config, err := jsonDecode(path)
		if err != nil {
			return err
		}
		editor = config.Editor
	}

	if editor == "" {
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

func getVinoteJsonPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, ".vinote.json"), nil
}

func existsConfigFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func jsonDecode(path string) (*Config, error) {
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

func askWhatUseEditor() string {
	var editor string
	fmt.Println("What do you use the editor?")
	fmt.Print("(vi or vim or nvim) > ")
	fmt.Scanln(&editor)
	return editor
}
