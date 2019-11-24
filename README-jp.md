# vinote

VimとNeoVim用のノートディレクトリへのアクセサツールです。

mattnさんの[memoアプリ](https://github.com/mattn/memo)のノート版を作ってみました！

> **NOTE** MacとLinuxのみ対応してます。

## Installation

Goをインストールして下記コマンドを実行すると簡単にインストールできます。

```
go get github.com/yasukotelin/vinote
```

## Setting

`.vinote.json` ファイルをホームディレクトリ（〜）に配置します。

その後下記を貼り付けて自分の設定を記載してください。

```json
{
    "path": "~/Dropbox/note",
    "editor": "nvim",
    "selector": "fzf"
}
```

| key      | value                                                                  | example        |
|----------|------------------------------------------------------------------------|----------------|
| path     | ノートのディレクトリを指定する                                         | ~/Dropbox/note |
| editor   | vimのパスを指定する（パスが通ってるなら普通にvimと指定すればいいです） | Vim, Neovim    |
| selector | fzfやpecoを指定すると `vinote select` コマンドで検索できます。         | fzf, peco      |

## Commands

| command       | value                                                   |
|---------------|---------------------------------------------------------|
| vinote        | どこからでも指定したノートディレクトリをVimで開けます！ |
| vinote open   | ノートディレクトリをFinderで開けます                    |
| vinote new    | 新しいファイルをどこからでも作成できます。              |
| vinote select | ノートディレクトリを対象にFuzzyFindできます。           |

## What is the Note direcotry?

ノートディレクトリは自分で適当に指定したディレクトリのことですが、下記のような形式でファイル管理することでノートのように扱えます。

```
note
├─fo
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

日付のディレクトリの中にindex.mdファイルを置くことで、添付の画像やPDFファイルなどを格納することができるので管理しやすいです。

また、 `vinote new` コマンドを使えば日付のディレクトリとindex.mdを作ってくれた上でVimを起動してくれるのでとても便利！

## Author

yasukotelin

## LICENSE

MIT
