package theme

import (
	"github.com/jesseduffield/lazygit/pkg/gocui"
	"github.com/jesseduffield/lazygit/pkg/gui/style"
)

type SemanticTheme struct {
	Focus                     style.TextStyle
	Error                     style.TextStyle
	InProgress                style.TextStyle
	Success                   style.TextStyle
	PrimaryAccent             style.TextStyle
	SecondaryAccent           style.TextStyle
	Base                      style.TextStyle
	Mantle                    style.TextStyle
	Text                      style.TextStyle
	ErrorBackground           style.TextStyle
	SuccessBackground         style.TextStyle
	PrimaryAccentBackground   style.TextStyle
	SecondaryAccentBackground style.TextStyle
	BaseForeground            style.TextStyle
	GocuiFocus                gocui.Attribute
	GocuiError                gocui.Attribute
	GocuiInProgress           gocui.Attribute
	GocuiSuccess              gocui.Attribute
	GocuiPrimaryAccent        gocui.Attribute
	GocuiSecondaryAccent      gocui.Attribute
	GocuiBase                 gocui.Attribute
	GocuiMantle               gocui.Attribute
	GocuiText                 gocui.Attribute
}

var Semantic = SemanticTheme{
	Focus:                     GetTextStyle([]string{"#cba6f7"}, false),
	Error:                     GetTextStyle([]string{"#f38ba8"}, false),
	InProgress:                GetTextStyle([]string{"#f9e2af"}, false),
	Success:                   GetTextStyle([]string{"#a6e3a1"}, false),
	PrimaryAccent:             GetTextStyle([]string{"#89b4fa"}, false),
	SecondaryAccent:           GetTextStyle([]string{"#f5c2e7"}, false),
	Base:                      GetTextStyle([]string{"#1e1e2e"}, true),
	Mantle:                    GetTextStyle([]string{"#181825"}, true),
	Text:                      GetTextStyle([]string{"#cdd6f4"}, false),
	ErrorBackground:           GetTextStyle([]string{"#f38ba8"}, true),
	SuccessBackground:         GetTextStyle([]string{"#a6e3a1"}, true),
	PrimaryAccentBackground:   GetTextStyle([]string{"#89b4fa"}, true),
	SecondaryAccentBackground: GetTextStyle([]string{"#f5c2e7"}, true),
	BaseForeground:            GetTextStyle([]string{"#1e1e2e"}, false),
	GocuiFocus:                GetGocuiStyle([]string{"#cba6f7"}),
	GocuiError:                GetGocuiStyle([]string{"#f38ba8"}),
	GocuiInProgress:           GetGocuiStyle([]string{"#f9e2af"}),
	GocuiSuccess:              GetGocuiStyle([]string{"#a6e3a1"}),
	GocuiPrimaryAccent:        GetGocuiStyle([]string{"#89b4fa"}),
	GocuiSecondaryAccent:      GetGocuiStyle([]string{"#f5c2e7"}),
	GocuiBase:                 GetGocuiStyle([]string{"#1e1e2e"}),
	GocuiMantle:               GetGocuiStyle([]string{"#181825"}),
	GocuiText:                 GetGocuiStyle([]string{"#cdd6f4"}),
}
