# gomi-rename

ゴミファイル名にリネームする CLI ツール

[![release](https://github.com/arrow2nd/gomi-rename/actions/workflows/release.yml/badge.svg?branch=main)](https://github.com/arrow2nd/gomi-rename/actions/workflows/release.yml)
[![test](https://github.com/arrow2nd/gomi-rename/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/arrow2nd/gomi-rename/actions/workflows/test.yml)

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

## Install

### Homebrew

```
brew tap arrow2nd/tap
brew install gomi-rename
```

### Scoop

```
scoop bucket add arrow2nd https://github.com/arrow2nd/scoop-bucket.git
scoop install arrow2nd/gomi-rename
```

### Binary

[Releases](https://github.com/arrow2nd/gomi-rename/releases) からお使いの環境にあったファイルをダウンロードしてください。

## Inspired by

> ゴミファイル名メーカー（@tooya2004 さん） | 診断メーカー
>
> https://shindanmaker.com/1119167
