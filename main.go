package main

import (
	"fmt"
	"os"

	"github.com/arrow2nd/gomi-rename/generator"
	"github.com/spf13/pflag"
)

type exitCode int

func (e exitCode) getInt() int {
	return int(e)
}

const (
	exitCodeOK exitCode = iota
	exitCodeParseErr
	exitCodeArgErr
	exitCodeGenErr
	exitCodeReadFileErr
	exitCodeRenameErr
)

var (
	appName = "gomi-rename"
	appDesc = "CLI tool to rename trashy file names"
	version = "develop"
	flags   = pflag.NewFlagSet(appName, pflag.ContinueOnError)
)

func init() {
	flags.IntP("number", "n", 5, "number of decorations to add to file name")
	flags.BoolP("dryrun", "d", false, "show execution results")
	flags.BoolP("version", "v", false, "show version")
	flags.BoolP("help", "h", false, "show help")
}

func main() {
	// パース
	if err := flags.Parse(os.Args[1:]); err != nil {
		exitErr(err.Error(), exitCodeParseErr)
	}

	// ヘルプ表示
	if help, _ := flags.GetBool("help"); help {
		exitHelp()
	}

	// バージョン表示
	if ver, _ := flags.GetBool("version"); ver {
		fmt.Printf("%s v.%s\n", appName, version)
		os.Exit(exitCodeOK.getInt())
	}

	// 引数不足
	if flags.NArg() <= 0 {
		exitErr("please specify one or more file names", exitCodeArgErr)
	}

	num, _ := flags.GetInt("number")
	isDryRun, _ := flags.GetBool("dryrun")

	for _, path := range flags.Args() {
		// ファイル名を生成
		newFileName, err := generator.Gen(path, num)
		if err != nil {
			exitErr(err.Error(), exitCodeGenErr)
		}

		fmt.Println(newFileName)

		// ドライラン
		if isDryRun {
			continue
		}

		// ファイルの存在確認
		if f, err := os.Stat(path); os.IsNotExist(err) || f.IsDir() {
			exitErr(
				fmt.Sprintf("file does not exist or is a directory: %s", path),
				exitCodeReadFileErr,
			)
		}

		// リネーム
		if err := os.Rename(path, newFileName); err != nil {
			exitErr(err.Error(), exitCodeRenameErr)
		}

		fmt.Println("  -> Done!")
	}
}

func exitErr(s string, c exitCode) {
	fmt.Fprintln(os.Stderr, s)
	os.Exit(c.getInt())
}

func exitHelp() {
	helpFormat := `%s

Usage:
  %s [flags] <filepath...>

Flags:
%s`

	fmt.Printf(helpFormat, appDesc, appName, flags.FlagUsages())
	os.Exit(exitCodeOK.getInt())
}
