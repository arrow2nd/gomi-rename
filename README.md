# gomi-rename

ゴミファイル名に自動でリネームする CLI ツール

[![release](https://github.com/arrow2nd/gomi-rename/actions/workflows/release.yml/badge.svg?branch=main)](https://github.com/arrow2nd/gomi-rename/actions/workflows/release.yml)
[![test](https://github.com/arrow2nd/gomi-rename/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/arrow2nd/gomi-rename/actions/workflows/test.yml)

## Usage

```
$ gomi-rename 最高の企画書.pdf
【優先度高】最高の企画書_修正版の修正版(6)_未記入アリ_最終確認版(6)_修正済(3)_訂正済.pdf
  -> Done!
```

```
$ gomi-rename -h
CLI tool for renaming to hellish file names

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
