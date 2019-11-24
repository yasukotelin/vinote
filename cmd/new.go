package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli"
	"github.com/yasukotelin/rlpath"
)

var scanner = bufio.NewScanner(os.Stdin)

func ExecNewCmd(c *cli.Context) error {
	config, err := ReadConfig()
	if err != nil {
		return err
	}
	configFullPath, err := config.GetFullPath()
	if err != nil {
		return err
	}

	scanner := &rlpath.Scanner{
		Prompt:  fmt.Sprintf("%s/", configFullPath),
		RootDir: configFullPath,
		OnlyDir: true,
	}

	scanPath, err := scanner.Scan()
	if err != nil {
		return err
	}

	title := readline("Title: ")
	dirName := formatToDirectoryName(title)

	createDir := filepath.Join(configFullPath, scanPath, dirName)

	for {
		in := readline(fmt.Sprintf("is this OK? (%s) [y/n] ", createDir))
		if in == "y" {
			break
		} else if in == "n" {
			return nil
		} else {
			fmt.Println("you should input the y or n")
		}
	}

	if err = os.MkdirAll(createDir, os.ModePerm); err != nil {
		return err
	}
	indexFile := filepath.Join(createDir, "index.md")
	file, err := os.Create(indexFile)
	if err != nil {
		return err
	}
	defer file.Close()

	if title != "" {
		h1Msg := fmt.Sprintf("# %s", title)
		_, err = fmt.Fprint(file, h1Msg)
		if err != nil {
			return err
		}
	}

	// Change the current directory to note directory.
	if err := os.Chdir(configFullPath); err != nil {
		return err
	}

	cmd := exec.Command(config.Editor, indexFile)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdin
	cmd.Stdin = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func readline(msg string) (line string) {
	fmt.Print(msg)
	if scanner.Scan() {
		line = scanner.Text()
	}
	return line
}

func formatToDirectoryName(s string) string {
	t := time.Now()
	if s == "" {
		return fmt.Sprintf("%s", t.Format("2006-01-02"))
	} else {
		return fmt.Sprintf("%s-%s", t.Format("2006-01-02"), strings.Join(strings.Split(s, " "), "-"))
	}
}
