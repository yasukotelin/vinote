package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
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
			// Action:  runConfig,
		},
		{
			Name:  "open",
			Usage: "opens the note directory with explorer(Finder)",
			// Action:  runOpen,
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
	// app.Action = runGonono

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
