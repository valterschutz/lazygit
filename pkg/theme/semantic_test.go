package theme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSemanticThemeUsesCatppuccinMocha(t *testing.T) {
	assert.Equal(t, GetTextStyle([]string{"#cba6f7"}, false), Semantic.Focus)
	assert.Equal(t, GetTextStyle([]string{"#f38ba8"}, false), Semantic.Error)
	assert.Equal(t, GetTextStyle([]string{"#f9e2af"}, false), Semantic.InProgress)
	assert.Equal(t, GetTextStyle([]string{"#a6e3a1"}, false), Semantic.Success)
	assert.Equal(t, GetTextStyle([]string{"#89b4fa"}, false), Semantic.PrimaryAccent)
	assert.Equal(t, GetTextStyle([]string{"#f5c2e7"}, false), Semantic.SecondaryAccent)
	assert.Equal(t, GetTextStyle([]string{"#1e1e2e"}, true), Semantic.Base)
	assert.Equal(t, GetTextStyle([]string{"#181825"}, true), Semantic.Mantle)
	assert.Equal(t, GetTextStyle([]string{"#cdd6f4"}, false), Semantic.Text)
	assert.Equal(t, GetTextStyle([]string{"#f38ba8"}, true), Semantic.ErrorBackground)
	assert.Equal(t, GetTextStyle([]string{"#a6e3a1"}, true), Semantic.SuccessBackground)
	assert.Equal(t, GetTextStyle([]string{"#89b4fa"}, true), Semantic.PrimaryAccentBackground)
	assert.Equal(t, GetTextStyle([]string{"#f5c2e7"}, true), Semantic.SecondaryAccentBackground)
	assert.Equal(t, GetTextStyle([]string{"#1e1e2e"}, false), Semantic.BaseForeground)
	assert.Equal(t, GetGocuiStyle([]string{"#cba6f7"}), Semantic.GocuiFocus)
	assert.Equal(t, GetGocuiStyle([]string{"#f38ba8"}), Semantic.GocuiError)
	assert.Equal(t, GetGocuiStyle([]string{"#f9e2af"}), Semantic.GocuiInProgress)
	assert.Equal(t, GetGocuiStyle([]string{"#a6e3a1"}), Semantic.GocuiSuccess)
	assert.Equal(t, GetGocuiStyle([]string{"#89b4fa"}), Semantic.GocuiPrimaryAccent)
	assert.Equal(t, GetGocuiStyle([]string{"#f5c2e7"}), Semantic.GocuiSecondaryAccent)
	assert.Equal(t, GetGocuiStyle([]string{"#1e1e2e"}), Semantic.GocuiBase)
	assert.Equal(t, GetGocuiStyle([]string{"#181825"}), Semantic.GocuiMantle)
	assert.Equal(t, GetGocuiStyle([]string{"#cdd6f4"}), Semantic.GocuiText)
}
