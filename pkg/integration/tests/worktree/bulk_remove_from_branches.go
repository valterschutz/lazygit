package worktree

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var BulkRemoveFromBranches = NewIntegrationTest(NewIntegrationTestArgs{
	Description: "Bulk branch cleanup skips dirty and locked worktrees, continues after failures, and confirms unmerged branches",
	SetupConfig: func(config *config.AppConfig) {
		config.GetUserConfig().Git.LocalBranchSortOrder = "alphabetical"
	},
	SetupRepo: func(shell *Shell) {
		shell.CloneIntoRemote("origin").EmptyCommit("initial commit")
		for _, name := range []string{"branch-01", "branch-02", "branch-03", "branch-04", "branch-05", "branch-06", "branch-07"} {
			shell.NewBranchFrom(name, "master").PushBranchAndSetUpstream("origin", name)
			if name == "branch-04" {
				shell.EmptyCommit("unmerged commit")
			}
			shell.Checkout("master")
			if name != "branch-02" {
				shell.AddWorktreeCheckout(name, "../"+name)
			}
		}
		shell.AddFileInWorktreeOrSubmodule("../branch-03", "untracked", "keep me")
		shell.RunCommand([]string{"git", "worktree", "lock", "../branch-05"})
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Branches().Focus().NavigateToLine(Contains("branch-01")).
			Press(keys.Universal.RangeSelectDown).
			Press(keys.Universal.RangeSelectDown).
			Press(keys.Universal.RangeSelectDown).
			Press(keys.Universal.RangeSelectDown).
			Press(keys.Universal.Remove)
		t.ExpectPopup().Menu().Title(Equals("Delete selected branches?")).Select(Contains("Delete local branches")).Confirm()
		t.ExpectPopup().Confirmation().Title(Equals("Delete selected branches?")).Content(Contains("Remove associated worktrees")).Confirm()
		t.ExpectPopup().Confirmation().Title(Equals("Force delete branch")).Content(Equals("These branches are not fully merged. Are you sure you want to delete them?\n\nbranch-04")).Cancel()
		// Cancelling the merge warning must not remove any worktrees or branches.
		t.Views().Branches().Lines(
			Contains("master"),
			Contains("branch-01 (worktree"),
			Contains("branch-02"),
			Contains("branch-03 (worktree"),
			Contains("branch-04 (worktree"),
			Contains("branch-05 (worktree"),
			Contains("branch-06 (worktree"),
			Contains("branch-07 (worktree"),
		).Press(keys.Universal.Remove)
		t.ExpectPopup().Menu().Title(Equals("Delete selected branches?")).Select(Contains("Delete local branches")).Confirm()
		t.ExpectPopup().Confirmation().Title(Equals("Delete selected branches?")).Content(Contains("Remove associated worktrees")).Confirm()
		t.ExpectPopup().Confirmation().Title(Equals("Force delete branch")).Content(Contains("branch-04")).Confirm()
		t.ExpectPopup().Alert().Title(Equals("Error")).Content(Contains("Cleanup completed with skipped or failed items")).
			Content(Contains("branch-03")).Content(Contains("branch-05")).Confirm()
		t.Views().Branches().Lines(
			Contains("master"),
			Contains("branch-03 (worktree"),
			Contains("branch-05 (worktree"),
			Contains("branch-06 (worktree"),
			Contains("branch-07 (worktree"),
		).NavigateToLine(Contains("branch-06")).Press(keys.Universal.RangeSelectDown).Press(keys.Universal.Remove)
		t.ExpectPopup().Menu().Title(Equals("Delete selected branches?")).Select(Contains("Delete local and remote branches")).Confirm()
		t.ExpectPopup().Confirmation().Title(Equals("Delete selected branches?")).Content(Contains("their upstream branches")).Confirm()
		t.Views().Branches().Lines(
			Contains("master"),
			Contains("branch-03 (worktree"),
			Contains("branch-05 (worktree"),
		)
		t.Views().Worktrees().Focus().Lines(
			Contains("(main worktree)"),
			Contains("branch-03"),
			Contains("branch-05"),
		)
		t.Views().Remotes().Focus().PressEnter()
		t.Views().RemoteBranches().Lines(
			Contains("branch-01"),
			Contains("branch-02"),
			Contains("branch-03"),
			Contains("branch-04"),
			Contains("branch-05"),
		)
	},
})
