# vinote

Quick access tool to a note directory with Vi, Vim, Neovim!

> **NOTE** Only supported MacOS or Linux.

## Installation

```
go get github.com/yasukotelin/vinote
```

## What is the Note direcotry?

If you want to write the note with your favorite editor, may be you have created the directory on local for markdown files.

This app is assuming that directory was created like the following.

```
note
├─food
│  └─2018-12-12-in-tokyo
│          index.md
│          picture1.jpeg
│          picture2.jpeg
│
├─game
│  ├─cod
│  │  └─2019-03-12-hint
│  │          index.md
│  │
│  └─ff
│      ├─2019-01-01-monster-data
│      │      index.md
│      │
│      └─2019-01-02-item-data
│              index.md
│
└─memo
    ├─2019-05-20-memo1
    │      index.md
    │
    └─2019-05-20-memo2
            index.md
```

## Setting

You have to put the `.vinote.json` on the `~` (home directory) or run `vinote config`.

```json
{
    "path": "~/Dropbox/note",
    "editor": "nvim",
    "selector": "fzf"
}
```

| key      | value                                     | example         |
|----------|-------------------------------------------|-----------------|
| path     | your note directory path                  | ~/Dropbox/note  |
| editor   | path to your vim (Gvim is not supported)  | vi, Vim, Neovim |
| selector | vinote can find with your setted selector | fzf, peco       |
