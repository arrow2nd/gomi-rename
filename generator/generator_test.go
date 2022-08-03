package generator_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/arrow2nd/gomi-rename/generator"
	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	// 正常
	files := []string{"テストファイル.png", "test.xml", "🐶.png"}

	for _, origin := range files {
		genName, err := generator.Gen(origin, 5)

		originExt := filepath.Ext(origin)
		originName := origin[:len(origin)-len(originExt)]

		assert.NoError(t, err, "正常に生成できるか")

		assert.True(
			t,
			strings.Contains(genName, originName),
			"指定した元ファイル名が残っているか",
			genName,
			originName,
		)

		assert.True(
			t,
			strings.Contains(genName, originExt),
			"指定した拡張子が残っているか",
			genName,
			originExt,
		)
	}

	// 異常
	_, err := generator.Gen("B.png", 10000)
	assert.ErrorContains(t, err, "too large value:", "生成数が上限を超えている場合、エラーが返るか")

	_, err = generator.Gen("C.jpeg", -1)
	assert.ErrorContains(t, err, "too small value:", "生成数が下限を下回っている場合、エラーが返るか")
}
