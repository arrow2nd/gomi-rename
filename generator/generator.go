package generator

import (
	"errors"
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

// Gen : ファイル名を生成
func Gen(path string, count int) (string, error) {
	// 生成数の上限・下限チェック
	if count < 0 {
		return "", errors.New("too small value: minimum value is 0")
	}

	if max := len(words); count > max {
		return "", fmt.Errorf("too large value: maximum value is %d", max)
	}

	rand.Seed(time.Now().UnixMicro())

	// リストをシャッフル
	listShuffle(titles)
	listShuffle(words)
	listShuffle(dateFormats)

	// パスからファイル名と拡張子を抽出
	ext := filepath.Ext(path)
	fileName := filepath.Base(path[:len(path)-len(ext)])

	// タイトルを付ける
	if titles[0] != "" {
		fileName = fmt.Sprintf("【%s】", titles[0]) + fileName
	}

	// 単語を付け足す
	addWords := []string{}
	for _, name := range words[:count] {
		// 40% の確率で 1~10 の数字を付ける
		if takeNum := rand.Intn(25) - 14; takeNum > 0 {
			name += fmt.Sprintf("(%d)", takeNum)
		}
		addWords = append(addWords, name)
	}

	fileName += "_" + strings.Join(addWords, "_")

	// 日付を置換
	fileName = strings.Replace(fileName, "YYYYMMDD", time.Now().Format(dateFormats[0]), 1)

	// パスを組み立てる
	return filepath.Join(filepath.Dir(path), fileName+ext), nil
}
