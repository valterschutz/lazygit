package worktree

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var BulkRemoveProtected = NewIntegrationTest(NewIntegrationTestArgs{
	Description: "Bulk branch cleanup cannot remove main or current worktrees",
	SetupConfig: func(config *config.AppConfig) {
		config.GetUserConfig().Git.LocalBranchSortOrder = "alphabetical"
	},
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("initial commit")
		shell.AddWorktree("master", "../current-worktree", "current-branch")
		shell.AddWorktree("master", "../old-worktree", "branch-01")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Worktrees().Focus().NavigateToLine(Contains("current-worktree")).Press(keys.Universal.Select)
		t.Views().Branches().Focus().Lines(
			Contains("current-branch").IsSelected(),
			Contains("branch-01"),
			Contains("master"),
		).Press(keys.Universal.RangeSelectDown).Press(keys.Universal.Remove)
		t.ExpectPopup().Menu().Title(Equals("Delete selected branches?")).Select(Contains("Delete local branches")).Confirm()
		t.ExpectToast(Contains("You cannot delete the checked out branch!"))
		t.ExpectPopup().Menu().Title(Equals("Delete selected branches?")).Cancel()
		t.Views().Branches().Press(keys.Universal.Return).NavigateToLine(Contains("branch-01")).Press(keys.Universal.RangeSelectDown).Press(keys.Universal.Remove)
		t.ExpectPopup().Menu().Title(Equals("Delete selected branches?")).Select(Contains("Delete local branches")).Confirm()
		t.ExpectPopup().Confirmation().Title(Equals("Delete selected branches?")).Content(Contains("Remove associated worktrees")).Confirm()
		t.ExpectPopup().Alert().Title(Equals("Error")).Content(Contains("master: Main and current worktrees cannot be removed.")).Confirm()
		t.Views().Branches().Lines(
			Contains("current-branch"),
			Contains("master"),
		)
		t.Views().Worktrees().Focus().Lines(
			Contains("current-worktree"),
			Contains("(main worktree)"),
		)
	},
})
