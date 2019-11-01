package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/yasukotelin/vinote/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "vinote"
	app.Version = "0.0.1"
	app.Usage = "Quick access tool to a note directory with Vi, Vim, Neovim!"
	app.Commands = []cli.Command{
		{
			Name:  "config",
			Usage: "creates the empty config file",
			Action: cmd.ExecConfigCmd,
		},
		{
			Name:  "open",
			Usage: "opens the note directory with explorer(Finder)",
			Action:  cmd.ExecOpenCmd,
		},
		{
			Name:    "new",
			Usage:   "creates the new note",
			Aliases: []string{"n"},
			// TODO newの挙動を再考
			// Action: runNew,
		},
		{
			Name:    "select",
			Usage:   "select a created note with selector command",
			Aliases: []string{"s"},
			// TODO 設定ファイルでセレクターを設定するような仕様にする
			// Action:  runFinder,
		},
	}

	app.Action = cmd.ExecVinoteCmd

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
