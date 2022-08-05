package generator

import "math/rand"

var (
	// ファイル名先頭に付くタイトル
	// 生成時に【】で囲まれます
	titles = []string{
		"",
		"重要",
		"未完成",
		"作業中",
		"社外秘",
		"要確認",
		"優先度高",
		"削除厳禁",
		"絶対消すな",
		"変更禁止",
		"上書き保存禁止",
		"社外持ち出し厳禁",
	}

	// ファイル名末尾に付くワード
	words = []string{
		"YYYYMMDD",
		"問い合わせ中",
		"未記入アリ",
		"没",
		"仮",
		"確定",
		"最新版",
		"最終版",
		"最終確認版",
		"確認終了版",
		"最終確認終了版",
		"修正版",
		"修正版の修正版",
		"訂正版",
		"修正済",
		"訂正済",
		"調整済",
		"反映済",
		"完成",
		"最終ver",
		"原本",
		"コピー",
		"バックアップ",
		"A案",
		"Final",
		"Final-Final",
		"Fix",
		"Copy",
		"Backup",
		"ver1",
	}

	// `YYYYMMDD` を日付に置換する際に使われるフォーマット文字列
	dateFormats = []string{
		"2006-01-02",
		"20060102",
		"060102",
		"20060102_1504",
		"060102_1504",
		"2006年01月02日",
		"1月2日",
		"1月分",
		"1月版",
		"1月2日変更済",
		"1月2日確認済",
	}
)

func listShuffle(l []string) {
	rand.Shuffle(len(l), func(i, j int) { l[i], l[j] = l[j], l[i] })
}
