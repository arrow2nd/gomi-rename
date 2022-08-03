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

	if max := len(names); count > max {
		return "", fmt.Errorf("too large value: maximum value is %d", max)
	}

	// 名前リストをシャッフル
	rand.Seed(time.Now().UnixMicro())
	rand.Shuffle(len(titles), func(i, j int) { titles[i], titles[j] = titles[j], titles[i] })
	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
	rand.Shuffle(len(dateFormats), func(i, j int) { dateFormats[i], dateFormats[j] = dateFormats[j], dateFormats[i] })

	// パスからファイル名と拡張子を抽出
	ext := filepath.Ext(path)
	fileName := filepath.Base(path[:len(path)-len(ext)])

	// タイトルを付ける
	if titles[0] != "" {
		fileName = fmt.Sprintf("【%s】", titles[0]) + fileName
	}

	// デコる
	decos := []string{}

	for _, name := range names[:count] {
		// 40% の確率で 1~10 のコピー回数を付ける
		if copyNum := rand.Intn(25) - 14; copyNum > 0 {
			name += fmt.Sprintf("(%d)", copyNum)
		}
		decos = append(decos, name)
	}

	fileName += "_" + strings.Join(decos, "_")

	// 日付を置換
	fileName = strings.Replace(fileName, "YYYYMMDD", time.Now().Format(dateFormats[0]), 1)

	// パスを組み立てる
	return filepath.Join(filepath.Dir(path), fileName+ext), nil
}
