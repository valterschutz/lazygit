# Repository Cleanup

The branch and worktree concepts used when cleaning up a repository.

## Language

**Branch**:
A named line of development. Removing its worktree does not delete the branch.

**Worktree**:
A working directory associated with a repository, with its own checked-out revision and local changes. Deleting a branch does not by itself remove its worktree.

**Dirty worktree**:
A worktree containing uncommitted changes or untracked files.

**Bulk cleanup**:
Removal of the selected branches and their associated worktrees. Items that cannot be safely removed are retained while the remaining items are processed.
