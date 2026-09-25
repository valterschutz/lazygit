package theme

import "github.com/jesseduffield/lazygit/pkg/gui/style"

type SemanticTheme struct {
	Focus           style.TextStyle
	Error           style.TextStyle
	InProgress      style.TextStyle
	Success         style.TextStyle
	PrimaryAccent   style.TextStyle
	SecondaryAccent style.TextStyle
	Base            style.TextStyle
	Mantle          style.TextStyle
	Text            style.TextStyle
}

var Semantic = SemanticTheme{
	Focus:           GetTextStyle([]string{"#cba6f7"}, false),
	Error:           GetTextStyle([]string{"#f38ba8"}, false),
	InProgress:      GetTextStyle([]string{"#f9e2af"}, false),
	Success:         GetTextStyle([]string{"#a6e3a1"}, false),
	PrimaryAccent:   GetTextStyle([]string{"#89b4fa"}, false),
	SecondaryAccent: GetTextStyle([]string{"#f5c2e7"}, false),
	Base:            GetTextStyle([]string{"#1e1e2e"}, true),
	Mantle:          GetTextStyle([]string{"#181825"}, true),
	Text:            GetTextStyle([]string{"#cdd6f4"}, false),
}
