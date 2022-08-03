# gomi-rename

ゴミファイル名にリネームする CLI ツール

## Usage

```
$ gomi-rename 最高の企画書.pdf
最高の企画書_未記入アリ_最終_こっちを使ってください_社外持ち出し厳禁_反映済み.pdf
```

```
$ gomi-rename -h
CLI tool to rename trashy file names

Usage:
  gomi-rename [flags] <filepath...>

Flags:
  -d, --dryrun       show execution results
  -h, --help         show help
  -n, --number int   number of decorations to add to file name (default 5)
  -v, --version      show version
```

## Inspired by

> ゴミファイル名メーカー（@tooya2004 さん） | 診断メーカー
> https://shindanmaker.com/1119167
