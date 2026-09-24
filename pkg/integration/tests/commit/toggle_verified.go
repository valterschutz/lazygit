package commit

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var ToggleVerified = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Mark commits as verified and unmark them again, both in the local commits and in the commits of another branch",
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
		shell.NewBranch("other")
		shell.EmptyCommit("three")
		shell.Checkout("master")
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

		t.Views().Branches().
			Focus().
			Lines(
				Contains("master").IsSelected(),
				Contains("other"),
			).
			SelectNextItem().
			PressEnter()

		t.Views().SubCommits().
			IsFocused().
			Lines(
				Contains("○").Contains("three").IsSelected(),
				Contains("○").Contains("two"),
				Contains("●").Contains("one"),
			).
			Press(keys.Commits.ToggleVerified).
			Lines(
				Contains("●").Contains("three").IsSelected(),
				Contains("○").Contains("two"),
				Contains("●").Contains("one"),
			).
			NavigateToLine(Contains("one")).
			Press(keys.Commits.ToggleVerified).
			Lines(
				Contains("●").Contains("three"),
				Contains("○").Contains("two"),
				Contains("○").Contains("one").IsSelected(),
			).
			PressEscape()

		t.Views().Commits().
			Lines(
				Contains("○").Contains("two"),
				Contains("○").Contains("one"),
			)
	},
})
