package commit

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var ToggleVerified = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Mark a commit as verified and unmark it again",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig: func(config *config.AppConfig) {
		// Relative to the test's repo, so that the file lands in the test's
		// own results directory next to it rather than in the shared $HOME.
		config.GetUserConfig().Gui.VerifiedCommitsFile = "../verified-commits"
	},
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("one")
		shell.EmptyCommit("two")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Commits().
			Focus().
			Lines(
				Contains("○").Contains("two").IsSelected(),
				Contains("○").Contains("one"),
			).
			Press(keys.Commits.ToggleVerified).
			Lines(
				Contains("●").Contains("two").IsSelected(),
				Contains("○").Contains("one"),
			).
			SelectNextItem().
			Press(keys.Commits.ToggleVerified).
			Lines(
				Contains("●").Contains("two"),
				Contains("●").Contains("one").IsSelected(),
			).
			SelectPreviousItem().
			Press(keys.Commits.ToggleVerified).
			Lines(
				Contains("○").Contains("two").IsSelected(),
				Contains("●").Contains("one"),
			)
	},
})
